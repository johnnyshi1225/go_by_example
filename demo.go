package main

import (
	. "fmt"
)

func main() {
	var b bool = true
	s := Sprintf("%t", b)
	Println(s)
}
