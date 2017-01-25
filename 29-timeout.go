package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case msg := <- c1:
		fmt.Println("Got message ",msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout occured1")
	}

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "result 2"
	}()

	select {
	case msg := <- c2:
		fmt.Println("Got message ", msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout occured2")
	}
}
