package main

import (
	"fmt"
)

func lopper(s string)  {
	for i := 0; i <= 3; i++ {
		fmt.Println("given string with counting ", s,i+1)
	}
}

func main() {

	lopper("first")

	go lopper("second: inside routine")

	go func(s string) {
		lopper(s)
	}("third: inside clousure routine")

	lopper("forth: all done")

	var input string
	fmt.Scanln(&input)
}
