package main

import (
	. "fmt"
)

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return r.width * 2 + r.height * 2
}

func (r *rect) setWidth(newWidth int) {
	r.width = newWidth
}

func main() {
	r := rect{2, 3}
	Println(r.area())

	r.setWidth(3)
	Println(r.area())
}
