// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"mitix-go/view_hlh/chapter29_ood/ood_main/encapsulation"
)


//用来测试，拿到对象以后如果是小写字段，对象中的字段可以获取么
func main() {
	mtx := encapsulation.NewMitix("hong", 20)
	fmt.Println(mtx.Name)
	fmt.Println(mtx.GetAge())

}
