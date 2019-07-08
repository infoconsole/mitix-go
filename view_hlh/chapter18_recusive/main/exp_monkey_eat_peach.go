// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	fmt.Println(eat_peach(1))
}

//递归方法  吃桃方法
func eat_peach(days int) int {
	if days <= 0 || days > 10 {
		panic("参数错误！！！")
	} else if (days == 10) {
		return 1
	} else {
		return (eat_peach(days+1) + 1) * 2
	}
}
