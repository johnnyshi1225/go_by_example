package main

import (
	. "fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 5

	zeroval(i)
	Println(i)

	zeroptr(&i)
	Println(i)
}
