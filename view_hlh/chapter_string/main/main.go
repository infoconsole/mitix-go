// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {

	//string类型的格式
	var str1 string = "infotech \n kaka"
	var str2 string = `infotech \n kaka`

	fmt.Println(str1)
	fmt.Println(str2)

	var poem string = `
						窗前明月光，
						疑似地上霜。
						举头望明月，
						低头思故乡。`
	fmt.Println(poem)
}
