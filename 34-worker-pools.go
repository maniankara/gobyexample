package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int)  {
	for j := range jobs {
		fmt.Println("Worker ", id, " starting job ", j)
		time.Sleep(time.Second)
		fmt.Println("Worker ", id, " finishing job ", j)
		results <- j * 2
	}
}

func main() {

	// declare jobs and results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// create 3 workers
	for i := 0; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// create 5 jobs for those 3 workers
	for j := 0; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 0; r <= 5; r++ {
		<- results
	}
}
