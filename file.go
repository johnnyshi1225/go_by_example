//#########################################################################
// Author: Johnny Shi
// Created Time: 2014-04-10 15:53:35
// File Name: file.go
// Description: 
//#########################################################################
package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("demo.go")
	if err != nil {
		return
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}

// vim: set noexpandtab ts=4 sts=4 sw=4 :
