package main

/**
	打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字
	立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次
	方＋5 的三次方＋3 的三次方。
 */

import (
	"fmt"
	"mitix-go/view-lnh/day2/example18/judge"
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
		go judge.DataJudge(i, piper, pipet)
	}
	for ; len(pipet) < m-n+1; {
		time.Sleep(100 * time.Millisecond)
	}
	var plen = len(piper)
	for i := 0; i < plen; i++ {
		t := <-piper
		fmt.Printf("%d\n", t)
	}
}
