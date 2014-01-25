package main

import "fmt"

const s string = "i am a constant"

func main() {
	fmt.Println(s)

	const n = 5000
	fmt.Println(float64(n))
}
