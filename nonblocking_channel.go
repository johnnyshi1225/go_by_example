package main

import (
	. "fmt"
	"time"
)

func main() {
	messages := make(chan string)

	select {
	case msg := <-messages:
		Println("Received msg:", msg)
	default:
		Println("No message in queue")
	}

	go func() {
		messages <- "some message"
	}()

	time.Sleep(time.Second)
	select {
	case msg := <-messages:
		Println("Received msg:", msg)
	default:
		Println("No message in queue")
	}
}
