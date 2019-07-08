package main

import (
	"fmt"
	ad "mitix-go/view_hlh/chapter15_function/function02_init/add"
)

func main() {
	//先进行打印
	fmt.Println(ad.Name)
	//调用一个方法
	ad.TestInit()
	//再进行打印
	fmt.Println(ad.Name)
	//打印常量
	fmt.Print(ad.Names)
}
