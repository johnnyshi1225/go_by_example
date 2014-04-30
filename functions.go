package main

import (
	. "fmt"
)

// Normal function
func add(a int, b int) int {
	return a + b
}

// Mutiple return values
func getError(code int) (int, string) {
	return code, "Page not found"
}

func xxx() (a, b, c, d string) {
	return "a1", "b2", "c3", "d4"
}

// Variadic function
func sum(nums ...int) int {
	Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	x, y := 2, 3
	Printf("%d + %d = %d\n", x, y, add(x, y))

	code, msg := getError(404)
	Printf("Error code: %d, error message: %s\n", code, msg)

	Println(sum(1, 2, 3))
	Println(sum(1, 2, 3, 4))

	nums := []int{1, 3, 5, 7, 9}
	Println(sum(nums...))

	Println(xxx())
}
