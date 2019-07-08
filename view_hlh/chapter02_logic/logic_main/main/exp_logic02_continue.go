// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {

	//continue案例
	//这里演示一下指定标签的形式来使用

	for i := 0; i < 4; i++ {
		//lable1: // 设置一个标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=", j)
		}
	}

here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i=", i, "j=", j)
		}
	}

}
