// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import "fmt"

func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string)
	smap := make(map[string]string, 16)
	smap["xx"] = "xxx"
	smap["yy"] = "yyy"
	smap["zz"] = "zzz"
	mapChan <- smap
	fmt.Println(mapChan)

}
