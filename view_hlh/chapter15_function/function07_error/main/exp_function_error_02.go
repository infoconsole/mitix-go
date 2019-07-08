// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type PathError struct {
	path    string
	op      string
	opTime  string
	message string
}

func (p *PathError) Error() string {
	return p.message + p.opTime
}

func main() {
	err := Open("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example10/main/sources1.txt")
	fmt.Println(err)
}

func Open(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &PathError{
			path:    fileName,
			op:      "read",
			opTime:  time.Now().Format("2006/01/02 15:04"),
			message: "this file is open fail",
		}
	}
	defer file.Close()
	if contents, err := ioutil.ReadAll(file); err == nil {
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file contents:", result)
	}
	return nil
}
