package main

import (
	. "fmt"
	"time"
)

func main() {
	// Process req per 200ms
	limiter := time.Tick(time.Millisecond * 200)
	reqs := make(chan int, 5)
	for i := 0; i < 5; i++ {
		reqs <- i
	}
	close(reqs)

	for req := range reqs {
		<-limiter
		Println("Process req:", req, time.Now())
	}

	// Allow 3 bursty reqs and then process req per 200ms 
	Println()
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	newReqs := make(chan int, 5)
	for i := 0; i < 5; i++ {
		newReqs <- i
	}
	close(newReqs)

	for req := range newReqs {
		<-burstyLimiter
		Println("Process req:", req, time.Now())
	}
}
