package main

import (
	"fmt"
	"math/rand"
)

func main() {
	expguessNumber()
}

//todo 还有问题  不知道怎么会执行3次
func expguessNumber() {
	var num = rand.Intn(100)

	for {
		var input int
		flag := false
		fmt.Scanf("%d\n", &input)
		switch {
		case input == num:
			fmt.Println("equal")
			flag = true
		case input > num:
			fmt.Println("bigger")
		case input < num:
			fmt.Println("less")
		default:
			fmt.Println("error")
		}

		if flag {
			break
		}
	}

}
