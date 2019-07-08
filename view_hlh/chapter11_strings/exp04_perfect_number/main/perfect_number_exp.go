// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
// 编程找出1000以内的所有完数
package main

import (
	"fmt"
	"strconv"
)

func main() {
	perfectNumber(1000)
}

func perfectNumber(n int) {
	for i := 1; i <= n; i++ {
		//加上1这个因子
		var sum = 0
		var strpf = ""
		for j := 1; j <= i/2; j++ {
			if j == 1 {
				sum = sum + j
				strpf = strpf + strconv.Itoa(j)
				continue
			}
			if i%j == 0 {
				sum = sum + j
				strpf = strpf + "+" + strconv.Itoa(j)
			}
		}
		if sum == i {
			fmt.Printf("%d = %s\n", i, strpf)
		}

	}

}
