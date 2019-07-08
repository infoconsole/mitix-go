package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("-----------------this is test exptimeFormat-----------------")
	exptimeFormat()
	fmt.Println("-----------------this is test exptimeFormat-----------------")

	fmt.Println("-----------------this is test exptimeUnix-----------------")
	exptimeUnix()
	fmt.Println("-----------------this is test exptimeUnix-----------------")
}

func exptimeFormat() {
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04"))

}

func exptimeUnix() {
	start := time.Now().UnixNano() / 1000
	test()
	end := time.Now().UnixNano() / 1000
	fmt.Printf("start - end is %d", end-start)
}

func test() {
	for i := 1; i < 1000; i++ {
		strconv.Itoa(i)
	}
}
