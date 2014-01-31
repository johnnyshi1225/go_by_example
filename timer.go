package main

import (
	. "fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		Println("Timer 2 stopped")
	}
}
