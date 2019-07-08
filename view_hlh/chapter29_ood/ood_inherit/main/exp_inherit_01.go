// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 继承

package main

import "fmt"

type Car struct {
	Name   string
	Weight int
}

func (this *Car) run() {
	fmt.Printf("this name is %s, and now running", this.Name)

}

type Bicycle struct {
	Car
	Wheel int
}

type Train struct {
	Car
}

func (this Train) String() string {
	str := fmt.Sprintf("name=[%s]  weight=[%d]", this.Name, this.Weight)
	fmt.Println(str)
	return str
}

func main() {
	var bike Bicycle
	bike.Name = "zixc"
	bike.Weight = 222
	bike.Wheel = 2
	fmt.Println(bike)
	bike.run()

	var train Train
	train.Name = "wuwuwu world!!"
	train.Weight = 100000
	//这里this实现的接口和指针实现的接口是不一样的
	fmt.Println(train)
}
