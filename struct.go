package main

import (
	. "fmt"
)

type ExtInfo struct {
	uid int
}

type Person struct {
	name string
	age int
	ext *ExtInfo
}

func main() {
	//Println(Person{"Johnny", 25})
	//Println(Person{name: "Johnny", age: 25})
	//Println(Person{age: 25, name: "Johnny"})

	//g := Person{"Grace", 25}
	//p := &g
	//Println(p.name)

	//p.age--
	//Println(g)
	g := &Person{"Grace", 25, &ExtInfo{123}}
	Println(g)
	h := g.ext
	Println(h)
	h.uid = 456
	Println(g.ext.uid)
}
