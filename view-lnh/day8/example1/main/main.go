// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// goroute计算阶乘

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]int)
	lock sync.Mutex
)

type tast struct {
	n int
}

func main() {
	for i := 0; i < 50; i++ {
		t := &tast{
			n: i,
		}
		go calc(t)
	}
	time.Sleep(time.Second * 10)
	for k, v := range m {
		fmt.Printf("key is %d ,and v is %d \n", k, v)
	}
}

func calc(t *tast) {
	sum := 1
	for i := 1; i < t.n; i++ {
		sum *= i
	}
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}
