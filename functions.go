package main

import (
	. "fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	x, y := 2, 3
	Printf("%d + %d = %d\n", x, y, add(x, y))
}
