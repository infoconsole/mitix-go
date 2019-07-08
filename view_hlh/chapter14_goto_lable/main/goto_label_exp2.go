// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// todo
package main

import (
	"fmt"
)

func main() {
	i := 0
HEAR:
	fmt.Printf("%d\n", i)
	i++
	if i == 5 {
		return
	}
	goto HEAR
}
