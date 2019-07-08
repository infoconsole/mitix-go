package swap

import "fmt"

//int 值传递  打印值以后
func Swap(a int, b int) (int, int) {
	fmt.Println("调用swap  进行值传递以后")
	fmt.Printf("in swap &a=%p,&b=%p\n", &a, &b)
	fmt.Println(" a b的值进行了交换  实际上是修改了a b 副本呢的值")
	tmp := a
	a = b
	b = tmp
	return a, b
}

func Swapx(a int, b int) (int, int) {
	fmt.Println("调用swapx  进行值传递以后")
	fmt.Printf("in a=%d,b=%d\n", a, b)
	fmt.Printf("in &a=%p,&b=%p\n", &a, &b)
	return b, a
}

/**
针对 值引用的方式
这个是不能进行交换的
原因，a,b两个变量都是副本的传输，所以是有问题的
*/
func Swap1(a *int, b *int) (*int, *int) {
	fmt.Println("调用swap1  进行地址传递")
	fmt.Printf("打印出来的a b 实际上是原来int的地址 a=%p,b=%p\n", a, b)
	//原来int的地址的地址
	fmt.Printf("in两个值是地址的地址 &a=%p,&b=%p\n", &a, &b)
	//交换了调用代码里面的地址，其实对于调用外面的地址没有变化
	tmp := a
	a = b
	b = tmp
	return a, b
}

func Swap2(a *int, b *int) (*int, *int) {
	fmt.Printf("in a=%d,b=%d\n", a, b)
	fmt.Printf("in &a=%x,&b=%x\n", &a, &b)
	//地址的值进行了交换，其实就是外面指向的地址的值   就是原来的值被改变了
	tmp := *a
	* a = * b
	* b = tmp
	return a, b
}
