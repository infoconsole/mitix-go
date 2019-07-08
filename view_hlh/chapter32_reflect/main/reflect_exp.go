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

func test(x interface{}) {

	var tye = reflect.TypeOf(x)
	fmt.Println(tye)

	v := reflect.ValueOf(x)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()

	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}else {
		fmt.Println(ok)
		fmt.Println(stu)
	}

}

type Student struct {
	Name  string
	Age   int
	Scope int
}

func main() {
	var a int = 200
	test(a)

	var s Student = Student{
		Name:  "info",
		Age:   100,
		Scope: 22,
	}

	test(s)


	//var x float64 = 3.4
	//fmt.Println("type:", reflect.TypeOf(x))
	//v := reflect.ValueOf(x)
	//fmt.Println("value:", v)
	//fmt.Println("type:", v.Type())
	//fmt.Println("kind:", v.Kind())
	//fmt.Println("value:", v.Float())
	//
	//fmt.Println(v.Interface())
	//fmt.Printf("value is %5.2e\n", v.Interface())
	//y := v.Interface().(float64)
	//fmt.Println(y)


}
