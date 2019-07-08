// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "time"

func main() {
	test_pipe()
}

func test_pipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	//如果管道深度是3  插入大于3个值会报死锁
	//因为这个进程是main进程
	//pipe <- 4

	println(len(pipe))
	var t int
	t = <-pipe
	println(t)
	time.Sleep(300)
}

