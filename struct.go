package main

import (
	. "fmt"
)

type Person struct {
	name string
	age int
}

func main() {
	Println(Person{"Johnny", 25})
	Println(Person{name: "Johnny", age: 25})
	Println(Person{age: 25, name: "Johnny"})

	g := Person{"Grace", 25}
	p := &g
	Println(p.name)

	p.age--
	Println(g)
}
