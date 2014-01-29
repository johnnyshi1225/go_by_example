package main

import (
	. "fmt"
	"time"
)

func basic() {
	msgChan := make(chan string)

	go func() {
		msgChan <- "ping"
	}()

	msg := <-msgChan
	Println(msg)
}

func buffered() {
	bufChan := make(chan string, 2)
	bufChan <- "some message"
	bufChan <- "another message"

	Println(<-bufChan)
	Println(<-bufChan)
}

func synchronization() {
	done := make(chan bool, 1)
	go func(done chan bool) {
		Println("Working...")
		time.Sleep(2 * time.Second)
		Println("Done")
		done <- true
	}(done)

	<-done
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "some message")
	pong(pings, pongs)
	Println(<-pongs)
}

func main() {
	// Basic
	basic()

	// Buffered channel
	buffered()

	// Synchronization
	synchronization()

	// Directions
	directions()
}
