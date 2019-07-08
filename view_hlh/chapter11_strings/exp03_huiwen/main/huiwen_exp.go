// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
// 左读完全相同的字符串。
package main

import (
	"fmt"
	"strings"
)

func main() {
	testHuiwen("长大大长")
}

func testHuiwen(source string) {
	chars := strings.Split(source, "")
	var lens = len(chars)

	for i := 0; i < lens/2; i++ {
		if chars[i] != chars[lens-1-i] {
			fmt.Println("this string is not huiwen")
		}
	}
	fmt.Println("this string is huiwen")

	//b := make([] byte, lens)
	//var bp = 0
	//for i := lens; i > 0; i -- {
	//	bp += copy(b[bp:], chars[i-1])
	//}
	//dest := string(b)
	//fmt.Printf("%v", dest)
}
