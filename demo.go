package main

import (
	"fmt"
)

type Hi interface {
	hi()
}

type Base struct {
	name string
}

func (b *Base) hi() {
	fmt.Println("i'm base")
}

type Child struct {
	Base
}

func (c *Child) hi() {
	fmt.Println("i'm child")
}

func main() {
	c := &Child{}
	c.hi()

	var hi Hi
	hi = c
	hi.hi()

	var b Base
	b = c
}
