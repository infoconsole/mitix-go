// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func add(a int, args ...int) int {
	var sum = a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func concat(a string, args ...string) (result string) {
	result = a
	for i := 0; i < len(args); i++ {
		result += args[i]
	}
	return result
}

func main() {
	sum := add(10, 6, 4, 3, 2, 1)
	fmt.Println(sum)

	result := concat("hello", "  ", "info", "  ", " world!!!")
	fmt.Println(result)
}
