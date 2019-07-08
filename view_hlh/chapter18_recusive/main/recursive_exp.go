// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 测试递归函数
package main

import (
	"fmt"
	"time"
)

func main() {
	var input int
	fmt.Scanf("%d\n", &input)
	expRecusive(input)
}

func expRecusive(i int) {
	fmt.Printf("this i is %d\n", i)
	time.Sleep(time.Second * 2)
	if i >= 0 {
		i--
		expRecusive(i)
	}
}
