// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func test0201(n int) {
	if n > 2 {
		n--
		test0201(n)
	}
	fmt.Println("n=", n)
}

func test22(n int) {
	if n > 2 {
		n-- //递归必须向退出递归条件逼进，否则就是无限循环调用
		test22(n)
	} else {
		fmt.Println("n=", n)
	}
}

func main() {

	//看一段代码
	//test0201(4) // ?通过分析来看下递归调用的特点
	test22(4) // ?通过分析来看下递归调用的特点
}
