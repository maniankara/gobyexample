package main

import (
	"fmt"
	"time"
)

func main()  {

	// simple switch
	switch time.Now().Month() {
	case time.January, time.December:
		fmt.Println("This month is either Jan or Dec")
	default:
		fmt.Println("This month is neither Jan nor Dec")

	}

	// Complex
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case string:
			fmt.Println("I am a string")
		case int:
			fmt.Println("I am an Int")
		default:
			fmt.Println("I am unknown", t)
		}
	}
	whatAmI(false)
	whatAmI(10)
	whatAmI("This is just a string")
	whatAmI(0.5)
}