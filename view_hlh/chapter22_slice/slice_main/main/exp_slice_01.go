// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	var exp_slice []int
	var array = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)

	//1. 切片初始化：var slice []int = arr[start:end],包含start到end之间的元素，但不包含end
	//2.  Var slice []int = arr[0:end]可以简写为 var slice []int=arr[:end]
	//3. Var slice []int = arr[start:len(arr)] 可以简写为 var slice[]int = arr[start:]
	//4. Var slice []int = arr[0, len(arr)] 可以简写为 var slice[]int = arr[:]
	//5. 如果要切片最后一个元素去掉，可以这么写： slice = slice[:len(slice)-1]

	//含义，前闭后开
	exp_slice = array[6:8]
	exp_slice[0] = 100
	fmt.Println(exp_slice)
	fmt.Println(len(exp_slice))
	// 这个是目前测试出来是去len(arr)-start,还不知道为什么
	fmt.Println(cap(exp_slice))
	fmt.Println(array)

	my_slice11 := make([]int, 6)
	fmt.Println(my_slice11)
	println(my_slice11)

	my_slice22 := make([]int, 3, 5)
	println(my_slice22) // [3/5]0xc42003df10
}
