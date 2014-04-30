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
	for i := 0; i < 5; i++ {
		Println(from, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	p := Person{name : "Johnny"}
	p.Friends = map[string]string{
		"w" : "15810294921",
		"x" : "14241823951",
	}
	go func() {
		//p.SayHi()
		if v, prs := p.Friends["w"]; prs {
			Println(v)
		}
	}()
	go func() {
		//p.SayHi()
		if v, prs := p.Friends["x"]; prs {
			Println(v)
		}
	}()

	/*
	foo("main")

	go foo("goroutine")

	go func(from string) {
		for i := 0; i < 5; i++ {
			Println(from, ":", i)
			time.Sleep(500 * time.Millisecond)
		}
	}("going")
	*/

	var s string
	Scanln(&s)
	Println("done")
}
