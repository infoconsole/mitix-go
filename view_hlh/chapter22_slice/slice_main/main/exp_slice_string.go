// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// string的切片操作
// 在go里面字符串是不可变的
package main

import "fmt"

func main() {
	testStringSlice()
	testStringSLice2()
	testStringSLice3()
	testStringSLice4()
}

func testStringSlice() {
	s := "Hello world!!!"
	s1 := s[0:5]
	s2 := s[6:]
	fmt.Println(s1)
	fmt.Println(s2)
}

//重要
func testStringSLice2() {
	s := "Hello world!!!"
	s1 := []rune(s)
	s1[1] = '咖'
	str := string(s1)
	fmt.Println(str)

}

func testStringSLice3() {
	s := "Hello world!!!"
	s1 := s[0:5]

	str := string(s1)
	fmt.Println(str)

}

func testStringSLice4() {
	s := "Hello world!!!"
	s1 := []byte(s)

	s1[1] = 'o'
	str := string(s1)
	fmt.Println(str)

}
