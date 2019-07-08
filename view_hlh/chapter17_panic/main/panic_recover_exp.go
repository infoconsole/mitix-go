// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 处理一个异常  panic recover
package main

import (
	"errors"
	"fmt"
)

func initConfig() (err error) {
	return errors.New("init config faild")
}

func main() {
	for {
		var input int
		fmt.Scanf("%d\n", &input)
		panicTest(input)
	}
}

func panicTest(in int) {
	// 手动进行panic
	err := initConfig()
	if err != nil {
		panic(err)
	}
	//返回执行错误
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error is-", err)
			//todo 业务逻辑
		}
	}()
	a := 100 / in
	fmt.Println(a)
}
