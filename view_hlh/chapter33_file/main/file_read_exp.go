// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// go语言中的文件读写

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Ioutil("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example5/main/testfile.txt")
	//OsIoutil("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example5/main/testfile.txt")
	//FileRead("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example5/main/testfile.txt")
	BufioRead("/Users/oldflame-jm/GoglandProjects/mitix-go/src/view-lnh/day7/example5/main/testfile.txt")
}

func Ioutil(name string) {
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
	result := strings.Replace(string(contents), "\n", "", 1)
	fmt.Println(result)
}

func OsIoutil(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("file open error: %v", err)
	}
	defer file.Close()
	if contents, err := ioutil.ReadAll(file); err == nil {
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file contents:", result)
	}
}

func FileRead(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("file open error: %v", err)
	}
	defer file.Close()
	//在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
	buf := make([]byte, 4096)
	if n, err := file.Read(buf); err == nil {
		fmt.Println("The number of bytes read:"+strconv.Itoa(n), "Buf length:"+strconv.Itoa(len(buf)))
		result := strings.Replace(string(buf), "\n", "", 1)
		fmt.Println("Use os.Open and File's Read method to read a file:", result)
	}
}

//file本身就是一个指针文件对象，被读取过一次就不能读取第二次
func BufioRead(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("file open error: %v", err)
	}
	defer file.Close()

	//一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
	reader := bufio.NewReader(file)
	//使用ReadString(delim byte)来读取delim以及之前的数据并返回相关的字符串.
	for {
		result, err := reader.ReadString(byte('\n'));
		if err == io.EOF {
			break
		}
		fmt.Println("使用ReadSlince相关方法读取内容:", result)
	}

	reader1 := bufio.NewReader(file)
	//注意:上述ReadString已经将buffer中的数据读取出来了，下面将不会输出内容
	//需要注意的是，因为是将文件内容读取到[]byte中，因此需要对大小进行一定的把控
	buf := make([]byte, 1024)
	//读取Reader对象中的内容到[]byte类型的buf中
	//上面已经读取了，下面直接提示EOF
	n, err := reader1.Read(buf);
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The number of bytes read:" + strconv.Itoa(n))
	fmt.Println("Use bufio.NewReader and os.Open read file contents to a []byte:", string(buf))

}
