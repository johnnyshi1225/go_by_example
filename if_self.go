package main

import "fmt"

func main() {
	num := 7

	if num % 2 == 0 {
		fmt.Printf("%d is even\n", num)
		fmt.Println(num, "is even")
	} else {
		fmt.Printf("%d is odd\n", num)
		fmt.Println(num, "is odd")
	}
}
