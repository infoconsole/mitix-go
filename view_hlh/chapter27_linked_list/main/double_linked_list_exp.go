// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 链表
//双向链表

package main

type Object interface{}

type DNode struct {
	data Object
	prev *DNode
	next *DNode
}

type DList struct {
	size uint64
	head *DNode
	tail *DNode
}

func (dList *DList) Init() {
	_dList := *(dList)
	_dList.size = 0   // 没车厢
	_dList.head = nil // 没车头
	_dList.tail = nil // 没车尾
}

func (dList *DList) Append(data Object) {
	newNode := new(DNode)
	(*newNode).data = data

	if (*dList).GetSize() == 0 { // 买个车头
		(*dList).head = newNode
		(*dList).tail = newNode
		(*newNode).prev = nil
		(*newNode).next = nil
	} else { //  挂在车队尾部
		(*newNode).prev = (*dList).tail
		(*newNode).next = nil
		(*((*dList).tail)).next = newNode
		(*dList).tail = newNode
	}

	(*dList).size++;
}
