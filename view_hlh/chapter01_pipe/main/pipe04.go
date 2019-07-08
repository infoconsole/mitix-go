package main

import (
	"fmt"
	"time"
)

/**
 *	使用go 和pipe进行异步计算返回
 */

func main() {

	cpipe := make(chan int, 1)

	go add3(1, 2, cpipe)

	sum1 := <-cpipe

	fmt.Println(sum1)
}

func add3(a int, b int, c chan int) int {
	sum := a + b
	time.Sleep(3 * time.Second)
	c <- sum
	return sum
}
