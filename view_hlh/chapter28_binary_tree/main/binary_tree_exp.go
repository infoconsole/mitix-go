// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 二叉树配置

package main

import "fmt"

type DataObject interface{}

type StudentTree struct {
	Name  string
	Age   int
	Data  DataObject
	Left  *StudentTree
	Right *StudentTree
}

func trans(root *StudentTree) {
	if root == nil {
		return
	}
	//前序遍历
	fmt.Println(root)
	trans(root.Left)
	//中序遍历
	//fmt.Println(root)
	trans(root.Right)
	//后序遍历
	//fmt.Println(root)
}

func main() {
	root := new(StudentTree)
	root.Name = "root"
	root.Age = 100

	left := new(StudentTree)
	left.Name = "left"
	left.Age = 99

	right := new(StudentTree)
	right.Name = "right"
	right.Age = 101

	root.Left = left
	root.Right = right

	trans(root)
}
