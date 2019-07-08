// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 链表
//每个节点包含下一个节点的地址，这样把所有的节点串起来了，通常把链表中的第一个节点叫做链表头

//如果有两个指针分别指向前一个节点和后一个节点，我们叫做双链表

package main

import (
	"fmt"
	"strconv"
	"sync"
)

var stu Student

//var stu = &Student{
//	Name: "hong.lvhang",
//	Age:  20,
//	Data: "info",
//	next: nil,
//}

type Date interface{}

type Student struct {
	mutex sync.Mutex
	Name  string
	Age   int
	Data  Date
	next  *Student
}

func init() {
	stu.Name = "hong.lvhang"
	stu.Age = 0
	stu.Data = "kaka"
}

func main() {
	fmt.Println(stu)

	for i := 1; i < 5; i++ {
		var p Student
		p.Name = "hong.lvhangxx" + strconv.Itoa(i)
		p.Age = i
		p.Data = "kaka1"
		p.next = nil
		//push
		fmt.Println(i)
		push(&p)
	}

	fmt.Println(stu)
	for stu.next != nil {
		fmt.Println(stu.next)
		stu = *stu.next
	}
}

func push(p *Student) {

	px := &stu
	for px.next != nil {
		px = px.next
	}
	px.mutex.Lock()
	px.next = p
	px.mutex.Unlock()
}
