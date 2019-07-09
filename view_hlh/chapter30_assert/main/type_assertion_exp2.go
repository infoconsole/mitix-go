// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 处理类型断言

package main

import "fmt"

func main() {
	var b int
	test(b)

	var x = "info"
	test(x)

}

func test(a interface{}) {
	c, ok := a.(int)
	if !ok {
		fmt.Println("convent failed!!")
		return
	}
	c += 3
	fmt.Println(c)
}
