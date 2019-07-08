// Copyright 2019, oldflame-jm. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 排序联系
package main

import "fmt"

func main() {
	//22,18,5,22,44,15,8,99,31,22
	var input1 [10]int
	fmt.Println("please input 10 number")
	fmt.Scanf("%d,%d,%d,%d,%d,%d,%d,%d,%d,%d\n", &input1[0], &input1[1], &input1[2], &input1[3], &input1[4], &input1[5], &input1[6], &input1[7], &input1[8], &input1[9])
	//冒泡排序
	//切面才能地址传递   如果是数组是值传递
	slice1 := input1[:]
	//testBubbleSort(slice1)
	testSelectionSort(slice1)
	//testInsertionSort(slice1)
	//testQuickSort(slice1)
	fmt.Println(slice1)

}

//冒泡排序
func testBubbleSort(numbers []int) {
	var sortLens = len(numbers)
	//比较len-1次   比较最大值为len -2
	for i := sortLens - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

//选择排序
func testSelectionSort(numbers []int) {
	var sortLens = len(numbers)
	for i := 0; i <= sortLens-1; i++ {
		var index = sortLens - 1
		var val = numbers[sortLens-1]
		for j := sortLens - 2; j >= i; j-- {
			if val > numbers[j] {
				index = j
				val = numbers[j]
			}
		}
		if i != index {
			numbers[i], numbers[index] = numbers[index], numbers[i]
		}
	}
}

//插入排序
func testInsertionSort(numbers []int) {
	var sortLens = len(numbers)
	//第0个为被比较对象
	for i := 1; i < sortLens; i++ {
		for j := i; j > 0; j-- {
			//类似于插队，一直往前插
			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			} else {
				//只要是小于了，就可以退出，因为从一开始前面就有序了
				break
			}
		}
	}
}

//快速排序
func testQuickSort(numbers []int) {
	quickSort(numbers, 0, len(numbers)-1)
}

func quickSort(numbers []int, start int, end int) {
	//从哪里开始
	stand := numbers[start]
	low := start
	high := end
	fmt.Println(numbers)
	fmt.Printf("stand %d,and start %d,and end %d\n", stand, low, high)
	//就是低位比高位大的，就退出了
	for ; low < high; {
		for ; low < high && numbers[high] >= stand; {
			//高位下降
			high--
		}
		//把高位复制到低位
		numbers[low] = numbers[high]
		for ; low < high && numbers[low] < stand; {
			//低位升
			low++
		}
		numbers[high] = numbers[low]
	}
	numbers[high] = stand

	//0到low-1  也就是  0到high-1
	if start < low {
		quickSort(numbers, start, low-1)
	}
	//空出high不用进行排序
	if high < end {
		quickSort(numbers, high+1, end)
	}
}
