package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

func main() {

	var readOps int64 = 0;
	var writeOps int64 = 0;

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	// read ops
	for r := 0; r < 100; r++ {
		total := 0
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddInt64(&readOps, 1)
			}
		}()
	}

	// write ops
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()

				atomic.AddInt64(&writeOps, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadInt64(&readOps)
	writeOpsFinal := atomic.LoadInt64(&writeOps)

	fmt.Println("read ops: ", readOpsFinal, " and write ops: ", writeOpsFinal)

	fmt.Println("state: ", state)
}
