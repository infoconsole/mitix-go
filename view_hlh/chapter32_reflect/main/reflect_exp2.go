// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"reflect"
)

type Skills interface {
	reading()
	running()
}

type Student1 struct {
	Name string
	Age  int
}

func (self *Student1) runing() {
	fmt.Printf("%s is running\n", self.Name)
}
func (self *Student1) reading() {
	fmt.Printf("%s is reading\n", self.Name)
}
func main() {
	stu1 := Student1{Name: "wd", Age: 22}
	inf := new(Skills)
	stu_type := reflect.TypeOf(stu1)
	inf_type := reflect.TypeOf(inf).Elem()                   // 特别说明，引用类型需要用Elem()获取指针所指的对象类型
	fmt.Println(stu_type.String())                           //main.Student
	fmt.Println(stu_type.Name())                             //Student
	fmt.Println(stu_type.PkgPath())                          //main
	fmt.Println(stu_type.Kind())                             //struct
	fmt.Println(stu_type.Size())                             //24
	fmt.Println(inf_type.NumMethod())                        //2
	fmt.Println(inf_type.Method(0), inf_type.Method(0).Name) // {reading main func() <invalid Value> 0} reading
	fmt.Println(inf_type.MethodByName("reading"))            //{reading main func() <invalid Value> 0} true

}
