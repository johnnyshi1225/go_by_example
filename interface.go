package main

import (
	. "fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return s.width * 2 + s.height * 2
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	Println(g)
	Printf("Area: %f\n", g.area())
	Printf("Perimeter: %f\n", g.perim())
}

func main() {
	s := square{2, 3}
	c := circle{4}

	measure(s)
	measure(c)
}
