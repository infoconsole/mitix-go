package main

/**
	对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
 */

import (
	"fmt"
	"mitix-go/view-lnh/day2/example19/judge"
	"time"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	piper := make(chan int, n+1)

	for i := 0; i <= n; i++ {
		go judge.MultiplySteps(i, piper)
	}
	for ; len(piper) < n+1; {
		time.Sleep(100 * time.Millisecond)
	}
	var total int
	var chainlen = len(piper)
	for i := 0; i < chainlen; i++ {
		t := <-piper
		total = total + t
	}
	fmt.Printf("%d\n", total)
}
