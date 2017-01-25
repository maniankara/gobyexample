package main

import (
	"time"
	"fmt"
)

func main() {
	ticker1 := time.NewTicker(time.Millisecond * 500)

	go func() {
		for i := range ticker1.C {
			fmt.Println("Ticking at ", i)
		}
	}()

	time.Sleep(time.Millisecond * 2002)
	ticker1.Stop()

	fmt.Println("Ticker stopped")

}

