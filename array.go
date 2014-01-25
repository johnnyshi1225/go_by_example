package main

import (
	. "fmt"
)

func main() {
	var a [5]int
	Println(a)

	a[3] = 52
	Println(a)
	Println(a[3])
	Println(len(a))

	b := [5]int{1, 3, 5, 7, 9}
	Println(b)
}
