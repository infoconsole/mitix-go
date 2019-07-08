// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 结构体

//1. 用来自定义复杂数据结构
//2. struct里面可以包含多个字段（属性）
//3. struct类型可以定义方法，注意和函数的区分
//4. struct类型是值类型
//5. struct类型可以嵌套
//6. Go语言没有class类型，只有struct类型
package main

import "fmt"

type Sxa struct {
	lad string
}

type Student struct {
	Sxa
	name string
	age  string
}

func main() {
	var stu1 Student
	stu1.name = "hong"
	stu1.age = "100"
	fmt.Println(stu1)
	fmt.Println(*&stu1)

	var stu2 = &Student{
		name: "dika",
		age:  "luobo80",
	}

	fmt.Println(*stu2)

	//var stu *Student =new(Student)
	//var stu *Student = &Student{}

	//stu.Sxa.lad = "info"

}
