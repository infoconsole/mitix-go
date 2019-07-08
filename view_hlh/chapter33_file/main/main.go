// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	Ccount     int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	file, err := os.Open("/Users/oldflame-jm/GoglandProjects/mitix-go/view-lnh/day7/example4/main/nxt.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var cc = &CharCount{}
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break

		}
		if err != nil {
			fmt.Printf("read string failed,err is %v", err)
			return
		}
		for _, c := range str {
			switch {
			case c >= 'a' && c <= 'z':
				cc.Ccount++
			case c >= 'A' && c <= 'Z':
				cc.Ccount++
			case c == ' ' || c == '\t':
				cc.SpaceCount++
			case c >= '0' && c <= '9':
				cc.NumCount++
			default:
				cc.OtherCount++
			}
		}
	}
	fmt.Println(*cc)
}
