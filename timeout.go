package main

import (
	. "fmt"
	"time"
)

const TIMEOUT = 1 * time.Second

func main() {
	c1 := make(chan string)
	go func() {
		Println("Working...")
		time.Sleep(2 * time.Second)
		c1 <- "result to c1"
	}()

	select {
	case ret := <-c1:
		Println(ret)
	case <-time.After(TIMEOUT):
		Printf("Timeout after %d s\n", TIMEOUT / time.Second)
	}
}
