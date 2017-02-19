package main

import (
	"fmt"
	"time"
)

func worker(done chan bool)  {
	fmt.Println("starting work")
	time.Sleep(time.Second * 2)
	fmt.Println("finishing work")
	done <- true
}

func main() {

	done := make(chan bool, 1)

	go worker(done)

	<- done // commenting out this will immediately exit with 0
}
