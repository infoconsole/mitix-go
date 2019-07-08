// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 斐波那契数
package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d\n", &input)
	ret := fab(input)
	fmt.Printf("the result of fab is %d", ret)
}

func fab(input int) int {
	if input <= 0 {
		panic("input can not less then zero")
	} else if input == 1 {
		return 1
	} else {
		return input + fab(input-1)
	}
}
