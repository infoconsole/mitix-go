// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"bytes"
	"fmt"
	"sync"
)

type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	UnLock()
}

type File interface {
	ReadWrite
	Lock
	Close()
}

type FileImpl struct {
	lock sync.Mutex
	Name string
}

func (file *FileImpl) Read(b bytes.Buffer) bool {
	file.lock.Lock()
	defer file.lock.Unlock()
	fmt.Printf("this is read file name is %s", file.Name)
	return true
}

func (file *FileImpl) Write(b bytes.Buffer) bool {
	panic("implement me")
}

func (file *FileImpl) Lock() {
	panic("implement me")
}

func (file *FileImpl) UnLock() {
	panic("implement me")
}

func (file *FileImpl) Close() {
	panic("implement me")
}

func main() {
	var file File
	var fileImpl FileImpl = FileImpl{
		Name: "Hong.Lvhang",
	}
	file = &fileImpl
	var buffer bytes.Buffer
	file.Read(buffer)
}
