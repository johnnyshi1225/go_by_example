package main

import (
	. "fmt"
	"time"
	//"runtime"
)

func foo(from string) {
	for i := 0; i < 5; i++ {
		Println(from, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	//runtime.GOMAXPROCS(2)

	foo("main")

	go foo("goroutine")

	go func(from string) {
		for i := 0; i < 5; i++ {
			Println(from, ":", i)
			time.Sleep(500 * time.Millisecond)
		}
	}("going")

	var s string
	Scanln(&s)
	Println("done")
}
