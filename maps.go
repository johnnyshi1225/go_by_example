package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]string)

	n := 10000000
	for i := 0; i < n; i++ {
		m[fmt.Sprintf("k%d", i)] = fmt.Sprintf("v%d", i)
	}

	fmt.Println("make data done")

	s := currentTimeMillis()
	fmt.Println(len(m))
	/*
	for k, _ := range m {
		delete(m, k)
	}
	fmt.Println(m)
	*/
	fmt.Printf("cost: %f ms\n", currentTimeMillis() - s)

}

func currentTimeMillis() float64 {
	return float64(time.Now().UnixNano()) / 1000000.0
}

