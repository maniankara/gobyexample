package main

import (
	"time"
	"fmt"
)

func main() {

	//
	// request limiter example
	requests := make(chan int, 5)

	for i := 0; i < 3; i++ {
		requests <- i
	}

	close(requests)

	// this channel gets tick every half second
	limiter := time.Tick(time.Millisecond * 500)

	for req := range requests {
		<- limiter // limit every half second
		fmt.Println("request", req, time.Now())
	}


	//
	// burst request limiter (each timeout can be varied)
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now() // also possible to insert time !
	}

	go func() {
		for t := range(time.Tick(time.Millisecond * 200)) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)

	for i := 0; i < 5; i++ {
		burstyRequest <- i
	}

	close(burstyRequest)

	for req := range burstyRequest {
		<- burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
