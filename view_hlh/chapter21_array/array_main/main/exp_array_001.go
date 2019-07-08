// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 数组
// 1. 数组：是同一种数据类型的固定长度的序列
// 2. 数组定义：var a [len]int，比如：var a[5]int  ,一旦定义，长度不能变
// 3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型
// 4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
// 5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
// 6. int类型数组默认情况下都会初始化成0
// 7. 数组是值类型，改变传入的值不会改变原值

// 切片

package main

import "fmt"

func main() {
	//testPanic()
	testArray()
	testArray2()

	//var age0 [5]int = [5]int{1, 2, 3}
	//var age1 = [5]int{1, 2, 3, 4, 5}
	//var age2 = [...]int{1, 2, 3, 4, 5}
	//var str = [5]string{3: "info", 4: "infotech"}
}

func testPanic() {
	var a [10]int
	j := 10
	a[0] = 100
	a[j] = 200
}

func testArray() {
	var a [10]int

	a[0] = 100
	a[1] = 111
	a[6] = 888

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for index, val := range a {
		fmt.Printf("a[%d]=%d  \n", index, val)
	}
}

func testArray2() {
	var a [3] int
	a[0] = 100

	a1 := a
	a1[0] = 200

	fmt.Printf("a[0] = %d\n", a[0])
}
