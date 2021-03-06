package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(msg string) {
		time.Sleep(time.Second * 1)
		c1 <- msg
	}("msg1")

	go func(msg string) {
		time.Sleep(time.Second * 2)
		c2 <- msg
	}("msg2")

	// selects which ever executes faster
	select {
	case msg1 := <- c1:
		fmt.Println("received ", msg1)
	case msg2 := <- c2:
		fmt.Println("received ", msg2)
	// default:
	// 	fmt.Println("nothing received")
	}
}
