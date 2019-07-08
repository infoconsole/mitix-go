// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 简单的http

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	err := http.ListenAndServe("0.0.0.0:9080", nil)
	checkErr(err, "http server start error")
}

func checkErr(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf(msg, err))
	}
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hello handler")
	writer.Write([]byte("this is an go http server"))
}
