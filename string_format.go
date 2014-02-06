package main

import (
	. "fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	Printf("%v\n", p)
	Printf("%+v\n", p)
	Printf("%#v\n", p)
	Printf("%T\n", p)

	Printf("%t\n", true)
	Println(true)

	Printf("%d\n", 123)
	Printf("%b\n", 123)
	Printf("%c\n", 103)
	Println("gg"[0])
	Printf("%x\n", 10010)

	Printf("%f\n", 3.14)
	Printf("%e\n", 314.15)
	Printf("%E\n", 314.15)

	Printf("%s\n", "\"string\"")
	Printf("%q\n", "string") // string with double-quote
	Printf("%x\n", "dex")

	Printf("%p\n", &p) // pointer

	s := Sprintf("a %s", "string")
	Println(s)

	Fprintf(os.Stderr, "an %s\n", "error")
}
