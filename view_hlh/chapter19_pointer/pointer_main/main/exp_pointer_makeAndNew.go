// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func expNew() {

	num1 := 100
	fmt.Printf("num1的类型%T , num1的值=%v , num1的地址%v\n", num1, num1, &num1)

	num2 := new(int) // *int
	//num2的类型%T => *int
	//num2的值 = 地址 0xc04204c098 （这个地址是系统分配）
	//num2的地址%v = 地址 0xc04206a020  (这个地址是系统分配)
	//num2指向的值 = 100
	*num2 = 100
	fmt.Printf("num2的类型%T , num2的值=%v , num2的地址%v\n num2这个指针，指向的值=%v",
		num2, num2, &num2, *num2)
}

func expMakeNew() {
	//新建指针
	s1 := new([]int)
	fmt.Println(s1)
	fmt.Println(*s1)

	//初始化空间，新建切片
	*s1 = make([]int, 5)
	(*s1)[0] = 100
	fmt.Println(*s1)

	//新建一个切片
	s2 := make([]int, 10)
	s2[0] = 100
	fmt.Println(s2)

}

func main() {
	expNew()
	expMakeNew()
}
