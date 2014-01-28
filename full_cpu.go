package main

import (
	. "fmt"
	"runtime"
)

func doNothing() {
	for {
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go doNothing()
	go doNothing()

	var s string
	Scanln(&s)
	Println("done")
}
