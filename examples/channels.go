package main

import "fmt"

func printer(s chan string)  {
	fmt.Println("String is: ", <-s)
}

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()


	var input string
	fmt.Scanln(&input)

	go func() {
		msg := <- messages

		fmt.Println(msg)
	}()

	fmt.Scanln(&input)
	// printer(messages)


}
