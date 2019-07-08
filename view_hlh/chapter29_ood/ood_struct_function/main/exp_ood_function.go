// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Golang中的方法是作用在特定类型的变量上，因此自定义类型，都可以有方法，而不仅仅是struct
// 定义：func (recevier type) methodName(参数列表)(返回值列表){}
package main

import "fmt"

func main() {
	var stu Student
	//(&stu).Init("hong.lvhang", 100)
	stu.Init("Hong.Lvhang", 200)
	fmt.Println(stu)

	var stu2 Student
	stu2.Init2("Info2", 200)
	fmt.Println(stu2)

	var ints integer
	ints = 1000
	ints.print()
	fmt.Println(ints)

	ints.set(2000)
	fmt.Println(ints)
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (stu *Student) Init(name string, age int) {
	stu.Name = name
	stu.Age = age
}

func (stu Student) Init2(name string, age int) {
	stu.Name = name
	stu.Age = age
}

type integer int

func (this integer) print() {
	fmt.Println("this is ", this)
}

func (pt *integer) set(val int) {
	*pt = integer(val)
}
