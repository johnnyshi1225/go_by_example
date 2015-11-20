package main

import (
	. "fmt"
	"time"
	"runtime"
)

type Person struct {
	name string
	Friends map[string]string
}

func (p Person) SayHi() {
	for {
		Println("Hi I'm", p.name)
		time.Sleep(1 * time.Second)
	}
}

func foo(from string) {
	time.Sleep(2000 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i <= 10000; i++ {
		go foo("goroutine")
		Printf("%d\n", runtime.NumGoroutine())
	}

	for {
		time.Sleep(100 * time.Millisecond)
		Printf("%d\n", runtime.NumGoroutine())
	}

	var s string
	Scanln(&s)
	Println("done")
}
