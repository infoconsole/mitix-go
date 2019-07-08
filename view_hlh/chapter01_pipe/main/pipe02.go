// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "time"

func main() {
	go test_pipe2()
	time.Sleep(5 * time.Second)
}

func test_pipe2() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	//如果这个异常是在goroute里面发生的，则不会有提示
	pipe <- 4
	println(len(pipe))
	var t int
	t = <-pipe
	println(t)
}
