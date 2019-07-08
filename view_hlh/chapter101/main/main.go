// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	crc32q := crc32.MakeTable(31)
	hc := crc32.Checksum([]byte("13818902375"), crc32q)
	fmt.Printf("%d", hc%4)
}
