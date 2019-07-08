// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	intChan := make(chan int, 10000)
	retChan := make(chan int, 10000)
	exitChan := make(chan string, 4)
	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for i := 0; i < 4; i++ {
		go calc(intChan, retChan, exitChan)
	}

	//等待所有协诚结束
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(retChan)
	}()

	for xv := range retChan {
		fmt.Println(xv)
	}

}

func calc(ori chan int, ret chan int, exitChan chan string) {
	var ss = true
	for v := range ori {
		for i := 2; i <= v/2; i++ {
			if v%i == 0 {
				ss = false
				break
			}
		}
		if !ss {
			ret <- v
		}
	}
	exitChan <- "true"
}
