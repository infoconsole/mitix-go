// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// sort排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	testSortInts()
	testSortStrings()
	testFloat64()
}

func testSortInts() {
	var a = [...]int{2, 66, 23, 48, 77, 12}
	//数组是值类型，需要传入切片
	sort.Ints(a[:])
	fmt.Println(a)
}

func testSortStrings() {
	var a = [...]string{"abc", "kaka", "infotdch", "old", "local", "zopp",}
	sort.Strings(a[:])
	fmt.Println(a)
}

func testFloat64() {
	var a = [...]float64{2.002, 66, 23.000, 48.455, 77.3333, 77.34}
	//数组是值类型，需要传入切片
	sort.Float64s(a[:])
	fmt.Println(a)

}

//sort.SearchInts(a []int, b int) 从数组a中查找b，前提是a必须有序
//sort.SearchFloats(a []float64, b float64) 从数组a中查找b，前提是a必须有序
//sort.SearchStrings(a []string, b string) 从数组a中查找b，前提是a必须有序
