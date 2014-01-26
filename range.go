package main

import (
	. "fmt"
)

func main() {
	a := []string{"foo", "bar", "baz", "qux"}
	for index, value := range a {
		Println(index, value)
	}
	for _, value := range a {
		Println(value)
	}

	m := map[string]string{
		"k1" : "v1",
		"k2" : "v2",
		"k3" : "v3",}
	for k, v := range m {
		Printf("%s -> %s\n", k, v)
	}

	for i, c := range "hello" {
		Println(i, c) // rune
	}
}
