// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	lensofletter, lensofblank, lensofnumber, lensofothers := testStatistics(input)
	fmt.Printf("now lensofletter=%d,lensofblank=%d,lensofnumber=%d,lensofothers=%d", lensofletter, lensofblank, lensofnumber, lensofothers)
}
func testStatistics(s string) (lensofletter, lensofblank, lensofnumber, lensofothers int) {
	t := []rune(s)
	for _, c := range t {
		switch {
		case unicode.IsLetter(c):
			lensofletter++
		case unicode.IsSpace(c):
			lensofblank++
		case unicode.IsNumber(c):
			lensofnumber++
		default:
			lensofothers++
		}
	}
	return
}
