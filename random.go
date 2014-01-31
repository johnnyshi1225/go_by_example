package main

import (
	. "fmt"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		Println(rand.Intn(10))
	}
}
