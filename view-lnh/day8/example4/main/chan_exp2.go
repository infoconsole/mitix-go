// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 100)

	for i := 0; i < 100; i++ {
		ch <- i
	}


	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}
