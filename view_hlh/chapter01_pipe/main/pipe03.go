package main

import (
	"fmt"
	"time"
)

/**
 *	使用go 和pipe进行异步计算返回
 */
var addpipe chan int

func main() {

	addpipe = make(chan int, 1)

	go add1(1, 2)

	sum1 := <-addpipe

	fmt.Println(sum1)
}

func add1(a int, b int) int {
	sum := a + b

	time.Sleep(5 * time.Second)

	addpipe <- sum
	return sum
}
