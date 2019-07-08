// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 模仿切片原理
package main

import "fmt"

type slice struct {
	ptr *[100]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func motify(s slice, index int, val int) {
	s.ptr[index] = val

}

func main() {
	var s1 slice
	s1 = make1(s1, 50)
	motify(s1, 10, 1000)
	fmt.Println(s1.ptr)

	var s2 []int = []int{1, 2, 3, 4, 5}
	motify1(s2)
	fmt.Println(s2)

}

func motify1(a []int) {
	a[1] = 1000
}
