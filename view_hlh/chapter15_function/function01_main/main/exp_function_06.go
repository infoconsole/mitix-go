// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

type add_func func(int, int) int

func add_f(a, b int) int {
	return a - b
}

func operator(op add_func, a int, b int) int {
	return op(a, b)
}

func main() {
	c := add_f
	fmt.Println(c)
	sum := operator(c, 10, 20)
	fmt.Println(sum)
}
