package main

import (
	. "fmt"
	"time"
)

func main() {
	m := make(map[string]string)
	m["k1"] = "v1"
	m["k2"] = "v2"
	Println(m)

	nm := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
		"k4": "v4",
	}

	Println(nm)

	Println(nm["k2"])

	_, prs := nm["k2"]
	Println(prs)
	delete(nm, "k2")
	Println(nm)
	Println(nm["k2"])
	_, prs = nm["k2"]
	Println(prs)

	s := currentTimeMillis()
	v, prs := nm["k1"]
	Printf("v = %s cost: %f ms\n", v, currentTimeMillis() - s)
}

func currentTimeMillis() float64 {
	return float64(time.Now().UnixNano()) / 1000000.0
}

