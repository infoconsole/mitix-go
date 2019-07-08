package main

import "fmt"

func main() {
	//直接执行外层循环，这里的疑问是跳槽以后 i的值为什么还在的
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is %d and j is %d\n", i, j)
		}
	}
}
