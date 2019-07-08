// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// io copy test

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example7/main/target.txt", "/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example7/main/sources.txt")
}

func CopyFile(dest string, src string) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer srcFile.Close()
	destFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer destFile.Close()
	io.Copy(destFile, srcFile)
}
