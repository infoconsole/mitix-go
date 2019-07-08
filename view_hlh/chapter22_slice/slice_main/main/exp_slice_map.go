// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

//切片和map一起使用

import "fmt"

func main() {
	var a []map[string]string
	a = make([]map[string]string, 10)
	a[0] = make(map[string]string)
	a[0]["key1"] = "val1"
	fmt.Println(a)
}
