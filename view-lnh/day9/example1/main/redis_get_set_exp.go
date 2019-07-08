// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 简单的redis  set get

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	stu := &student{
		Id:    100000,
		Name:  "hong.lvhang",
		Email: "hlh19850220@163.com",
	}

	data, err := json.Marshal(stu)
	//对本次链接进行set操作
	hash := sha256.New()
	hash.Write(data)
	result := hash.Sum(nil)
	hashId := hex.EncodeToString(result)
	fmt.Println(hashId)
	//连接
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	checkErr(err, "redis conn error")
	defer c.Close()

	//设置值
	_, setErr := c.Do("set", hashId, data)
	checkErr(setErr, "redis command error")

	//使用redis的string类型获取set的k/v信息
	r, getErr := redis.String(c.Do("get", hashId))
	checkErr(getErr, "redis command error")
	fmt.Println(r)
}

func checkErr(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf(msg, err))
	}
}
