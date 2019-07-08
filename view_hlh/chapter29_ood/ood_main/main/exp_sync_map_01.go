// Copyright 2019, oldflame-jm. All rights reserved.

//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// todo
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwLock sync.RWMutex

func main() {
	//testModify()
	testModify2()
}

func testModify2() {
	var count int32
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3

	for i := 0; i < 2; i++ {
		go func(b map[string]int) {
			for {
				rwLock.Lock()
				b["a"] = rand.Intn(100)
				rwLock.Unlock()
			}
		}(a)
	}

	for i := 0; i < 10; i++ {
		go func(b map[string]int) {
			for {
				rwLock.RLock()
				atomic.AddInt32(&count, 1)
				//fmt.Printf("gorouter %d,val %d", i, b["a"])
				//fmt.Println(b["a"])
				rwLock.RUnlock()

			}
		}(a)
	}

	time.Sleep(time.Second * 10)
	fmt.Println(count)

}

func testModify() {
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3

	for i := 0; i < 2; i++ {
		go func(b map[string]int) {
			for {
				lock.Lock()
				b["a"] = rand.Intn(100)
				fmt.Println(b["a"])
				lock.Unlock()
			}
		}(a)
	}

	time.Sleep(time.Second * 10)
}
