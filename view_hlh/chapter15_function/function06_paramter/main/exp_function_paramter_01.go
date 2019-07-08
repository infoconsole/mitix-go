// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	var a = 100
	var b = make(chan int, 100)
	fmt.Println(a)
	//值类型  int float bool string 数组  struct
	//引用类型  指针  slice  map  chan
	fmt.Printf("a的指针地址,%p \n", &a)
	fmt.Printf("b是一个chan,%p \n", b)

	modify1(a)
	fmt.Println(a)

	modify2(&a)
	fmt.Println(a)
}

func modify1(a int) {
	a = 10
	return
}
func modify2(a *int) {
	*a = 10
}
