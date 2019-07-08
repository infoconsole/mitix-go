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

func read(ch chan int) {
	for i := 0; i < 100; i++ {
		x := <-ch
		fmt.Println(x)
	}
}

func wirte(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func main() {
	intChan := make(chan int, 10)

	go read(intChan)

	go wirte(intChan)

	time.Sleep(time.Second * 10)
}
