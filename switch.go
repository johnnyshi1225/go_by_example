package main

import (
	. "fmt"
	"time"
)

func main() {
	current := time.Now()
	Println(current.Weekday())

	switch current.Weekday() {
	case time.Saturday, time.Sunday:
		Println("It's the weekend")
	default:
		Println("It's the weekday")
	}

	Println(current.Hour())
	switch {
	case current.Hour() < 12:
		Println("It's before noon")
	default:
		Println("It's after noon")
	}
}
