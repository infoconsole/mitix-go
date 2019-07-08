// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 创建切面的几种方式
//var slice []type = make([]type, len)
//slice  := make([]type, len)
//slice  := make([]type, len, cap)

package main

import "fmt"

func main() {
	testCreate1()
	testCreate2()
	testCreate3()
	testCreate4()

}

func testCreate4() {
	var s1 = []int{1, 2, 3}
	var s2 = []int{4, 5, 6}
	var s3 = append(s1, s2...)
	fmt.Println(s3)

	copy(s1, s2)
	fmt.Println(s1)

	for index, val := range s1 {
		fmt.Printf("this key val is [%d=%d]\n", index, val)
	}

	var s4 = []int{1, 2, 3, 4, 5}
	s5 := s4[0:2]
	//resize
	s5 = s5[0:5]
	fmt.Println(s5)

}

func testCreate3() {
	var slice []int = make([]int, 10)
	fmt.Println(slice)
	fmt.Println(&slice[0])
}

func testCreate2() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9}
	var slice = arr[2:8]
	fmt.Println(&arr[2])
	fmt.Println(&slice[0])
	fmt.Println(slice[0])
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	fmt.Println(&slice)

}

func testCreate1() {
	var slice = make([]int, 2, 4)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	fmt.Println(slice)
	fmt.Println(cap(slice))

}
