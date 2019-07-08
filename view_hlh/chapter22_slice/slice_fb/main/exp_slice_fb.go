// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 使用非递归的方式实现呢斐波那契数
package main

import "fmt"

func main() {
	var input int
	fmt.Scanf("%d\n", &input)
	testFab(input)
}

func testFab(num int) {
	var array []int
	array = make([]int, num)
	for i := 0; i < num; i++ {
		if i > 1 {
			array[i] = array[i-1] + array[i-2]
		} else {
			array[i] = 1
		}
		fmt.Printf("array[%d]=%d\n", i, array[i])
	}
	//var sum int
	//for i := 0; i < num; i++ {
	//	sum += array[i]
	//}
	//fmt.Printf("sum is [%d]\n", sum)
}
