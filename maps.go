package main

import (
	. "fmt"
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
		"k4": "v4",}

	Println(nm)

	Println(nm["k2"])

	_, prs := nm["k2"]
	Println(prs)
	delete(nm, "k2")
	Println(nm)
	Println(nm["k2"])
	_, prs = nm["k2"]
	Println(prs)

}
