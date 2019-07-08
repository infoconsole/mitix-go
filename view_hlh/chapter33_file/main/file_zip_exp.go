// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	ReadZipFile("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example6/main/ltt.zip")
}

func ReadZipFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	//这里读取压缩文件有问题
	fz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(fz)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%v", line)
	}
}
