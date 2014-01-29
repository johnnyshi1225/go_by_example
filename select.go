/*
Execute this like "time go run select.go" and the final execution time is as much as 2 second.
*/
package main

import (
	. "fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i <2; i++ {
		select {
		case msg1 := <-c1:
			Println("Received:", msg1)
		case msg2 := <-c2:
			Println("Received:", msg2)
		}
	}
}
