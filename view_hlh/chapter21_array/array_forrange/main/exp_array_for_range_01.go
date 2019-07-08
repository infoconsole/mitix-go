// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 二维数组
package main

import "fmt"

func main() {
	var f [2][3]int = [...][3]int{{2, 5, 6}, {8, 7, 11}}
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d,%d)=%d\n", k1, k2, v2)
		}
	}
}
