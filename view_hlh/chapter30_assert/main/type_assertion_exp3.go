// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//

package main

import "fmt"

type nn struct {
	Name string
}

func main() {
	var b int
	classifier(b, nn{}, "12345", 22.22222, true)
}

//通过反射去动态获取类型
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n ", i)
		case float64:
			fmt.Printf("param #%d is a float64\n ", i)
		case int, int64:
			fmt.Printf("param #%d is a int\n ", i)
		case nil:
			fmt.Printf("param #%d is a nil\n ", i)
		case string:
			fmt.Printf("param #%d is a string\n ", i)
		default:
			fmt.Printf("param #%d is a unknow\n ", i)
		}
	}
}
