// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 双向链表实现

package main

import "fmt"

type List interface {
	//从最后插入
	Add(data interface{}) bool
	//定点插入
	AddIndex(index int, data interface{}) bool
	//获取不删除
	Get(index int) interface{}
	//获取并删除
	Remove(index int) interface{}
	//队列深度
	Size() int
}

type Deque interface {
	//增加，从队首
	AddFirst(data interface{}) bool
	//增加，从队尾
	AddLast(data interface{}) bool
	//取出，从队首
	RemoveFirst() interface{}
	//取出，从队尾
	RemoveLast() interface{}
	//查看，从队首
	Peek() interface{}
	//查看，从队首
	PeekFirst() interface{}
	//查看，从队尾
	PeekLast() interface{}
	//入栈
	Push(data interface{})
	//出栈
	Pop() interface{}
}

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

type LinkedList struct {
	size int
	head *Node
	tail *Node
}

func (list *LinkedList) Init() {
	_list := *list
	_list.size = 0
	_list.head = nil
	_list.tail = nil
}

//断开关联
func (list *LinkedList) unLink(node *Node) interface{} {
	data := node.data
	prev := node.prev
	next := node.next
	//关联链表  清除前节点
	if prev == nil {
		list.head = next
	} else {
		prev.next = next
		node.prev = nil
	}
	//清除后节点
	if next == nil {
		list.tail = prev
	} else {
		next.prev = prev
		node.next = nil
	}

	//清除数据
	node.data = nil
	list.size --
	return data
}

//移除头
func (list *LinkedList) unLinkFirst(h *Node) interface{} {
	data := h.data
	next := h.next
	h.data = nil
	h.next = nil
	list.head = next

	if next == nil {
		list.tail = nil
	} else {
		next.prev = nil
	}
	list.size--
	return data
}

//解绑尾部节点
func (list *LinkedList) unLinkLast(t *Node) interface{} {
	data := t.data
	prev := t.prev
	t.data = nil
	t.prev = nil
	list.tail = prev
	if prev == nil {
		list.head = nil
	} else {
		prev.next = nil
	}
	list.size--
	return data
}

func (list *LinkedList) linkLast(data interface{}) {
	t := list.tail
	//新建最后的节点
	node := Node{
		data: data,
		prev: t,
	}
	list.tail = &node
	if t == nil {
		list.head = &node
	} else {
		t.next = &node
	}
	list.size++
}

func (list *LinkedList) linkFirst(data interface{}) {
	h := list.head
	//创建节点
	node := Node{
		data: data,
		next: h,
	}
	list.head = &node
	if h == nil {
		list.tail = &node
	} else {
		h.prev = &node
	}
	list.size++
}

func (list *LinkedList) linkBefore(data interface{}, succ *Node) {
	pred := succ.prev
	node := Node{
		data: data,
		prev: pred,
		next: succ,
	}
	succ.prev = &node
	if pred == nil {
		list.head = &node
	} else {
		pred.next = &node
	}
	list.size++
}

func (list *LinkedList) checkElementIndex(index int) {
	if !(index >= 0 && index < list.size) {
		panic(fmt.Sprintf("index %d out of Bounds", index))
	}
}

func (list *LinkedList) checkPositionIndex(index int) {
	if !(index >= 0 && index <= list.size) {
		panic(fmt.Sprintf("index %d out of Bounds", index))
	}
}

func (list *LinkedList) Add(data interface{}) bool {
	list.linkLast(data)
	return true
}

func (list *LinkedList) node(index int) *Node {
	if index < (list.size >> 1) {
		x := list.head
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x
	} else {
		x := list.tail
		for i := list.size - 1; i > index; i-- {
			x = x.prev
		}
		return x
	}
}

func (list *LinkedList) AddIndex(index int, data interface{}) bool {
	list.checkPositionIndex(index)
	if index == list.size {
		list.linkLast(data)
	} else {
		list.linkBefore(data, list.node(index))
	}
	return true
}

func (list *LinkedList) AddFirst(data interface{}) bool {
	list.linkFirst(data)
	return true
}

func (list *LinkedList) AddLast(data interface{}) bool {
	list.linkLast(data)
	return true
}

func (list *LinkedList) Get(index int) interface{} {
	list.checkElementIndex(index)
	return list.node(index).data
}

func (list *LinkedList) Remove(index int) interface{} {
	list.checkElementIndex(index)
	return list.unLink(list.node(index))
}

func (list *LinkedList) RemoveFirst() interface{} {
	return list.unLinkFirst(list.head)
}

func (list *LinkedList) RemoveLast() interface{} {
	return list.unLinkLast(list.tail)
}

func (list *LinkedList) Peek() interface{} {
	return list.head.data
}

func (list *LinkedList) PeekFirst() interface{} {
	return list.head.data
}

func (list *LinkedList) PeekLast() interface{} {
	return list.tail.data
}

func (list *LinkedList) Push(data interface{}) {
	list.linkLast(data)
}

func (list *LinkedList) Pop() interface{} {
	return list.unLinkLast(list.tail)
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) String() string {
	x := list.head
	var str string
	str = x.data.(string)
	for i := 0; i < list.size-1; i++ {
		x = x.next
		str = str + x.data.(string)
	}
	return str
}

func main() {
	var linkList LinkedList
	var list List
	list = &linkList
	list.Add("info")
	fmt.Println(list)

}
