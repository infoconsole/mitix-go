// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 定时器

package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println(v)
	}
}
