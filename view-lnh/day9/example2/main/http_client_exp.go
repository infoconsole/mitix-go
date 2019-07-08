// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// todo

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://www.baidu.com")
	//checkErr(err, "http connect err")

	data, _ := ioutil.ReadAll(res.Body)
	//checkErr(err, "http response err")

	fmt.Println(string(data))
}
