package main

/**
判断 101-200 之间有多少个素数，并输出所有素数
*/

import (
	"fmt"
	"mitix-go/view-lnh/day2/example17/judge"
	"time"
)

func main() {
	var n int
	var m int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	piper := make(chan int, m-n+1)
	pipet := make(chan int, m-n+1)
	for i := n; i <= m; i++ {
		go judge.PrimeJupdge(i, piper, pipet)
	}
	for ; len(pipet) < m-n+1; {
		time.Sleep(100 * time.Millisecond)
	}
	for i := 0; i < len(piper); i++ {
		t := <-piper
		fmt.Printf("%d\n", t)
	}
}
