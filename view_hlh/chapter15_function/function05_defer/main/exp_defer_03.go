// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// defer的价值在于执行完毕以后肯定会调用以后释放资源

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fileRead(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("file open error: %v", err)
	}
	//调用defer释放资源
	defer file.Close()
	//在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
	buf := make([]byte, 4096)
	if n, err := file.Read(buf); err == nil {
		fmt.Println("The number of bytes read:"+strconv.Itoa(n), "Buf length:"+strconv.Itoa(len(buf)))
		result := strings.Replace(string(buf), "\n", "", 1)
		fmt.Println("Use os.Open and File's Read method to read a file:", result)
	}
}

func main() {
	fileRead("/aaa.txt")
}
