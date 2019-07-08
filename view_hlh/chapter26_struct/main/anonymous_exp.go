// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 结构体的匿名字段

package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

type mc chan int

type Train struct {
	Car
	int
	start time.Time
}

func main() {
	var t Train
	t.int = 200
	t.Name = "info"
	t.Age = 200
	fmt.Println(t)
}
