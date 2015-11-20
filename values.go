package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("7.0 / 3.0 =", 7.0 / 3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var x bool
	fmt.Println(x)

	var cid interface{}
	cid = 1232
	/*
	cidStr, ok := cid.(string)
	if !ok {
		fmt.Println("not ok")
	}
	*/
	cidStr := fmt.Sprint(cid)
	fmt.Println(cidStr)
}
