//展示了接口的继承方式
package main

import (
	"fmt"
)

type BBInterface interface {
	test01()
}

type CCInterface interface {
	test02()
}

type AAInterface interface {
	BBInterface
	CCInterface
	test03()
}

//如果需要实现AInterface,就需要将BInterface CInterface的方法都实现
type Stuu struct {
}

func (stu Stuu) test01() {

}
func (stu Stuu) test02() {

}
func (stu Stuu) test03() {

}

type T interface {
}

func main() {
	var stu Stuu
	var a AAInterface = stu
	a.test01()

	var t T = stu //ok
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}
