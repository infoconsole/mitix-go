// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// todo
package main

import "fmt"

func main() {
	testMap1()
	testMap2()

}


//var a map[string]string
// var a map[string]int
// var a map[int]string
// var a map[string]map[string]string
func testMap2() {
	amap := make(map[string]map[string]string, 100)

	//map的赋值操作，要先处理map的初始化
	amap["key1"] = make(map[string]string)
	amap["key1"]["key1"] = "val1"
	amap["key1"]["key2"] = "val2"
	amap["key1"]["key3"] = "val3"
	amap["key1"]["key4"] = "val4"
	amap["key1"]["key5"] = "val5"
	fmt.Println(amap)

	//map的取值
	val, ok := amap["key2"]
	if ok {
		fmt.Println(val)
	} else {
		amap["key2"] = make(map[string]string)
		amap["key2"]["kaka"] = "info"
	}
	fmt.Println(amap)

	//遍历
	for k1, v1 := range amap {
		for k2, v2 := range v1 {
			fmt.Printf("this map is %s,%s=%s\n", k1, k2, v2)
		}
	}
	fmt.Println("-----------------delete----------------------")
	//删除一个存在的key
	delete(amap, "key2")
	//删除一个不存在的key
	delete(amap, "key3")
	fmt.Println("-----------------delete----------------------")

	//遍历
	for k1, v1 := range amap {
		for k2, v2 := range v1 {
			fmt.Printf("this map is %s,%s=%s\n", k1, k2, v2)
		}
	}

	//len
	fmt.Println(len(amap))

}

//var map1 map[keytype]valuetype
//var a map[string]string
//var a map[string]int
//var a map[int]string
//var a map[string]map[string]string
func testMap1() {
	//申明
	var a = map[string]string{"hello": "world1"}
	fmt.Println(a)

	//赋内存
	a = make(map[string]string, 10)
	fmt.Println(a)

	//put操作
	a["hello"] = "world2"
	fmt.Println(a)

	//delete(a, “hello”)
	//
	//len(a)

}
