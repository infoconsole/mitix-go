// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"mitix-go/ooad-go/simplefactory/exp02/vehicle"
)

//可以证明两个对象地址相同  实现了单例模式
func main() {
	pCar := vehicle.NewCar()
	fmt.Printf("%p", pCar)
	fmt.Println(pCar)

	pCar2 := vehicle.NewCar()
	fmt.Printf("%p", pCar2)
	fmt.Println(pCar2)

}
