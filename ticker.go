package main

import (
	. "fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 3)
	ticker.Stop()
	Println("Ticker stopped")
}
