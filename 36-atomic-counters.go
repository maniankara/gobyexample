package main

import (
	"sync/atomic"
	"time"
	"fmt"
)

func main() {

	var ops int64 = 0;

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddInt64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()

	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)

	fmt.Println("Value: ", opsFinal)

}
