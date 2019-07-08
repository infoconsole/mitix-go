// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// map遍历是有序还是无序的，在go中  map是无序的

package main

import (
	"fmt"
	"sort"
)

func main() {
	testMapSort()
	fmt.Println("--------------sort-------------")
	testMapSort2()
}

//无序代码
func testMapSort() {
	var a map[int]int
	a = make(map[int]int)
	a[1] = 6
	a[8] = 9
	a[3] = 10
	a[12] = 88
	a[5] = 55
	a[44] = 18
	for k, v := range a {
		fmt.Printf("kv %d=%d\n", k, v)
	}
}

//有序代码
func testMapSort2() {
	var a map[int]int
	a = make(map[int]int)
	a[1] = 6
	a[8] = 9
	a[3] = 10
	a[12] = 88
	a[5] = 55
	a[44] = 18

	var keys []int
	for k, _ := range a {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		fmt.Printf("kv %d=%d\n", v, a[v])
	}
}
