// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	var a interface{}
	fmt.Printf("%d %T\n", a, a)
	var b int
	a = b
	c := a.(int)
	fmt.Printf("%d %T\n", a, a)
	fmt.Printf("%d %T\n", c, c)
}

