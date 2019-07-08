// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 文件写入

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	name := "/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example5/main/testwirte.txt"
	content := "Hello, xxbandy.github.io!\n"
	//WriteWithIoutil(name, content)
	//WriteWithFileWrite(name, content)
	//WriteWithIo(name, content)
	WriteWithBufio(name, content)
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("写入文件成功:", content)
	}
}


//const (
//	O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
//	O_WRONLY int = syscall.O_WRONLY // 只写打开文件
//	O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
//	O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
//	O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
//	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
//	O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
//	O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
//)

//使用os.OpenFile()相关函数打开文件对象，并使用文件对象的相关方法进行文件写入操作
func WriteWithFileWrite(name,content string){
	fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	defer fileObj.Close()
	if _,err := fileObj.WriteString(content);err == nil {
		fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.",content)
	}
	contents := []byte(content)
	if _,err := fileObj.Write(contents);err == nil {
		fmt.Println("Successful writing to thr file with os.OpenFile and *File.Write method.",content)
	}
}


//使用io.WriteString()函数进行数据的写入
func WriteWithIo(name,content string) {
	fileObj,err := os.OpenFile(name,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		fmt.Println("Failed to open the file",err.Error())
		os.Exit(2)
	}
	if  _,err := io.WriteString(fileObj,content);err == nil {
		fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.",content)
	}
}


//使用bufio包中Writer对象的相关方法进行数据的写入
func WriteWithBufio(name,content string) {
	if fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644);err == nil {
		defer fileObj.Close()
		writeObj := bufio.NewWriterSize(fileObj,5)
		//
		if _,err := writeObj.WriteString(content);err == nil {
			fmt.Println("Successful appending buffer and flush to file with bufio's Writer obj WriteString method",content)
		}

		//使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
		buf := []byte(content)
		if _,err := writeObj.Write(buf);err == nil {
			fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Write method.",content)
			if  err := writeObj.Flush(); err != nil {panic(err)}
			fmt.Println("Successful flush the buffer data to file ",content)
		}
	}
}