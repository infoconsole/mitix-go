// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go语言接口定义

package main

import (
	"fmt"
)

type Test interface {
	Print()
	Sleep()
}

type People struct {
	Name string
	Age  int
	//types.Slice
}

type Student struct {
	Name  string
	Age   int
	Score int
}

func (this Student) Sleep() {
	panic("implement me")
}

func (this People) Sleep() {
	fmt.Println(this.Age)
	//fmt.Println(this.Slice)
}

func (this Student) Print() {
	fmt.Println(this.Name)
	fmt.Println(this.Age)
	fmt.Println(this.Score)
}

func (this People) Print() {
	fmt.Println(this.Age)
}

//noinspection GoBinaryAndUnaryExpressionTypesCompatibility
func main() {
	var t Test
	var stu Student
	stu.Age = 100
	stu.Name = "hong"

	t = stu
	t.Print()
	//t.Sleep()

	var t1 Test
	var people People
	people.Name = "hong.lvhang"
	people.Age = 20000
	//people.Slice = make([]int, 10)

	t1 = people
	t1.Sleep()
	t1.Print()
}
