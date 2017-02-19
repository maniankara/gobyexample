package main

import (
	"time"
	"fmt"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<- timer1.C
	fmt.Println("Timer1 has expired !")

	timer2 := time.NewTimer(time.Second * 4)
	go func() { // functions waits for 4 seconds for expiry
		<- timer2.C
		fmt.Println("Timer2 has expired !")
	}()

	timer2.Stop()
	fmt.Println("Timer2 stopped")


}
