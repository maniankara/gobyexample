package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j,more := <- jobs // more is always true until close is called
			if more {
				fmt.Println("job received ", j)
			} else {
				fmt.Println("all jobs done")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
	}

	close(jobs)

	fmt.Println("All jobs done successfully")

	<- done
}
