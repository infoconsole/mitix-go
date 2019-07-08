// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"mitix-go/view_hlh/chapter15_function/function06_paramter/swap"
)

func main() {
	var a = 100
	var b = 200
	//打印初始化的地址信息
	fmt.Println("初始化的a,b的地址信息")
	fmt.Printf("now &a=%p,&b=%p\n", &a, &b)
	fmt.Printf("-----------------------\n")

	swap.Swap(a, b)
	fmt.Printf("swap交换以后 a=%d,b=%d\n", a, b)
	fmt.Printf("-----------------------\n")

	swap.Swapx(a, b)
	fmt.Printf("now a=%d,b=%d\n", a, b)
	fmt.Printf("-----------------------\n")

	swap.Swap1(&a, &b)
	fmt.Printf("now a=%d,b=%d\n", a, b)
	fmt.Printf("-----------------------\n")
	swap.Swap2(&a, &b)
	fmt.Printf("now a=%d,b=%d\n", a, b)

}
