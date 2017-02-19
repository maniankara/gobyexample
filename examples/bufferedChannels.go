package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	go func() {
		messages <- "message1"
		messages <- "message2"
	}()

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
