// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	var xChan chan int
	xChan = make(chan int, 10)
	xChan <- 10
	fmt.Println(xChan)
}
