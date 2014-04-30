package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 7

	if num % 2 == 0 {
		fmt.Printf("%d is even\n", num)
		fmt.Println(num, "is even")
	} else {
		fmt.Printf("%d is odd\n", num)
		fmt.Println(num, "is odd")
	}

	s := "12313x"
	i, err := strconv.Atoi(s)
	fmt.Println(i, err)

}
