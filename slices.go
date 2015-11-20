package main

import (
	. "fmt"
)

type Person struct {
	Name string
}

func hehe(list []interface{}) {
	for _, l := range list {
		Println(l)
	}
}

func main() {
	/*
	s := make([]string, 3)
	Println(s)

	s[0] = "first"
	s[1] = "test"
	s[2] = "demo"
	Println(s)
	Println(len(s))

	s = append(s, "hehe")
	Println(s)
	Println(len(s))

	c := make([]string, len(s))
	copy(c, s)
	Println(c)
	Println(len(c))

	Println(s[0:2])
	Println(s[2:])
	Println(s[:3])

	l := make([]interface{}, 0)
	ps := make([]*Person, 0)

	ps = append(ps, &Person{Name : "xxx"})
	ps = append(ps, &Person{Name : "yyy"})

	l = append(l, &Person{Name : "johnny"})
	l = append(l, &Person{Name : "grace"})

	Println(s)
	s = nil
	Println(s)
	for _, i := range s {
		Println(i)
	}
	*/
	l := []int{}
	l = append(l, 1)
	Println(l)
}
