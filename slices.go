package main

import (
	. "fmt"
)

func main() {
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

}
