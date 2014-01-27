package main

import (
	. "fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	Println(nextInt())
	Println(nextInt())
	Println(nextInt())
	Println(nextInt())

	Println()
	newSeq := intSeq()
	Println(newSeq())
	Println(newSeq())
}
