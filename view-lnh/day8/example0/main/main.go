// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"time"
)

func main() {
	go test()
	var i int
	for {
		fmt.Printf("this i is %d\n", i)
		time.Sleep(time.Second * 2)
	}
}

func test() {
	var i int
	for {
		fmt.Printf("test i is %d\n", i)
		time.Sleep(time.Second)
	}
}
