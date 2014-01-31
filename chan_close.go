package main

import (
	. "fmt"
)

func main() {
	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	close(queue)

	for {
		e, more := <-queue
		if (more) {
			Println("Receive", e)
		} else {
			Println("No more")
			break
		}
	}

	// Also I can range over this channel like this
	for elem := range queue {
		Println(elem)
	}
}
