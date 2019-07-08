// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// struct里面对于字段可以有一个解释

package main

import (
	"encoding/json"
	"fmt"
)

type Studentt struct {
	Name  string `json:"name""`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	var stut Studentt = Studentt{
		Name:  "Hong.Lvhang",
		Age:   100,
		Score: 3,
	}
	data, err := json.Marshal(stut)
	if err != nil {
		fmt.Println("json encode faild,err ", err)
		return
	}
	fmt.Println(string(data))

}
