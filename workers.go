package main

import (
	. "fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		Println("Worker", id, "process job:", i)
		time.Sleep(time.Second)
		results <- i * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < 9; j++ {
		jobs <- j
	}
	close(jobs)

	for j := 0; j < 9; j++ {
		Println(<-results)
	}
}
