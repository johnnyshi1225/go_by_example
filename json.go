//#########################################################################
// Author: Johnny Shi
// Created Time: 2015-08-12 14:54:57
// File Name: json.go
// Description: 
//#########################################################################
package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

type ExtInfo struct {
	Logo string
	AppName string
}

/*
type ExtInfo struct {
	Logo string	`json:"logo"`
	AppName string	`json:"app_name"`
}
*/

func main() {
	//var extInfo ExtInfo
	data := `{"logo":"1015960","name":"\u5ea6\u5468\u672b","app_name":"\u5ea6\u5468\u672b\uff0d\u5468\u672b\u6e38\uff0c\u4e0b\u5348\u8336\uff0c\u81ea\u52a9\u9910\uff0c\u7f8e\u597d\u751f\u6d3b\u5c3d\u5728\u5ea6\u5468\u672b","text":"\u4e0b\u8f7d"}`
	//var d map[string]interface{}
	d := &ExtInfo{}
	err := json.Unmarshal([]byte(data), &d)
	fmt.Println(err)
	fmt.Println(d.Logo, d.AppName)

	dec := json.NewDecoder(strings.NewReader(data))
	err = dec.Decode(&d)
	fmt.Println(err)
	fmt.Println(d.Logo, d.AppName)
}


// vim: set noexpandtab ts=4 sts=4 sw=4 :
