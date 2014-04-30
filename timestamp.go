//#########################################################################
// Author: Johnny Shi
// Created Time: 2014-03-27 10:34:49
// File Name: timestamp.go
// Description: 
//#########################################################################

package main

import (
	"fmt"
	"time"
)

func main() {
	s := currentTimeMillis()
	fmt.Printf("%f ms\n", currentTimeMillis() - s)

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("15"))
}

func currentTimeMillis() float64 {
	return float64(time.Now().UnixNano()) / 1000000.0
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
