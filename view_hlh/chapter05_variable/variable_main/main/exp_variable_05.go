// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

var a = "G"

func main() {
	n()
	mm()
	fmt.Println(a)
	x()
	fmt.Println(a)
}

func n() {
	fmt.Println(a)
}

func mm() {
	a := "OO"
	fmt.Println(a)
}

func x() {
	a = "XX"
	fmt.Println(a)
}
