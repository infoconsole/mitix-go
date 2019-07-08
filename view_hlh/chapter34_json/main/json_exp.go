// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// json处理

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type List struct {
	Users []User
	Dtss  string
	Lad   int
	Amap  map[string]map[string]string
}

func main() {
	amap := make(map[string]map[string]string, 100)

	//map的赋值操作，要先处理map的初始化
	amap["key1"] = make(map[string]string)
	amap["key1"]["key1"] = "val1"
	amap["key1"]["key2"] = "val2"
	amap["key1"]["key3"] = "val3"
	amap["key1"]["key4"] = "val4"
	amap["key1"]["key5"] = "val5"
	fmt.Println(amap)

	list := &List{}
	user := &User{
		UserName: "oldflame-Jm",
		Age:      18,
		Birthday: "2001-01-22",
		Sex:      "女",
		Email:    "vev_sss@163.com",
		Phone:    "13777728883",
	}
	list.Users = append(list.Users, *user)
	list.Users = append(list.Users, User{
		UserName: "aaa",
		Age:      22,
		Birthday: "2001-01-22",
		Sex:      "XX",
		Email:    "vev_sss@163.com",
		Phone:    "13777728883",
	})
	list.Dtss = "infotech"
	list.Lad = 100
	list.Amap = amap
	data, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	var list1 List
	json.Unmarshal(data, &list1)
	fmt.Println(list1)
}
