// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package log

import (
	"fmt"
	"io"
	"os"
)

func ClientLog(content string) {
	name := "/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/chitchat/client.log"
	writeLog(name, content)
}

func ServerLog(content string) {
	name := "/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/chitchat/server.log"
	writeLog(name, content)
}

//使用io.WriteString()函数进行数据的写入
func writeLog(name, content string) {
	fileObj, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("log: Failed to open the file", err.Error())
		os.Exit(2)
	}
	if _, err := io.WriteString(fileObj, content); err != nil {
		fmt.Println("log: Fail appending to the file with os.OpenFile and io.WriteString.", content)
	}
}
