package main

import "fmt"

func main() {
	var a bool = true
	var b int32 = 12345
	var c byte = 'c'
	var d string = "info"
	var f float32 = 88234.887

	//%v表示  相应值的默认格式。在打印结构体时，“加号”标记（%+v）会添加字段名
	fmt.Println("-------------------print bool with %v--------------------")
	fmt.Printf("相应值的默认格式 %v\n", a)
	fmt.Printf("相应值的默认格式 %v\n", b)
	fmt.Printf("相应值的默认格式 %v\n", c)
	fmt.Printf("相应值的默认格式 %v\n", d)
	//浮点数打印出来也不准
	fmt.Printf("相应值的默认格式 %v\n", f)
	fmt.Println("-------------------print bool with %v--------------------")

	//相应值的Go语法表示
	fmt.Println("-------------------print bool with %#v--------------------")
	fmt.Printf("相应值的Go语法表示 %#v\n", a)
	fmt.Printf("相应值的Go语法表示 %#v\n", b)
	//byte类型的go语法表示为16进制
	fmt.Printf("相应值的Go语法表示 %#v\n", c)
	fmt.Printf("相应值的Go语法表示 %#v\n", d)
	fmt.Println("-------------------print bool with %#v--------------------")

	//相应值类型的go语法表示
	fmt.Println("-------------------print bool with %T--------------------")
	fmt.Printf("T go语言类型 %T\n", a)
	fmt.Printf("T go语言类型 %T\n", b)
	fmt.Printf("T go语言类型 %T\n", c)
	fmt.Printf("T go语言类型 %T\n", d)
	fmt.Println("-------------------print bool with %T--------------------")

	//打印单词 true或者false
	fmt.Println("-------------------print bool with bool --------------------")
	fmt.Printf("t打印单词true或者false  %t\n", a)
	fmt.Println("-------------------print bool with bool --------------------")

	fmt.Println("-------------------print bool with int --------------------")
	//二进制表示binary
	fmt.Printf("b是二进制binary  %b\n", b)
	//相应Unicode码点所表示的字符
	b=110
	fmt.Printf("c是Unicode   %c\n", b)
	//十进制
	fmt.Printf("d十进制   %d\n", b)
	//八进制
	fmt.Printf("o八进制   %o\n", b)
	//单引号围绕的字符字面值
	fmt.Printf("单引号围绕的字符 %q\n", b)
	fmt.Printf("十六进制小写 %x\n", b)
	fmt.Printf("十六进制大写 %X\n", b)
	fmt.Printf("Unicode格式 %U\n", b)
	fmt.Println("-------------------print bool with int --------------------")

	fmt.Println("-------------------print bool with float --------------------")
	fmt.Printf("无小数部分的，指数为二的幂的科学计数法 %b\n", f)
	fmt.Printf("科学计数 %e\n", f)
	fmt.Printf("科学计数 %E\n", f)
	fmt.Printf("有小数点而无指数 %f\n", f)
	fmt.Printf("%g\n", f)
	fmt.Printf("%G\n", f)

	fmt.Println("-------------------print bool with float --------------------")

	fmt.Println("-------------------print bool with string --------------------")
	fmt.Printf("字符串或切片的无解译字节 %s\n", d)
	fmt.Printf("双引号围绕的字符串，由Go语法安全地转义 %q\n", d)
	fmt.Printf("十六进制，小写字母，每字节两个字符 %x\n", d)
	fmt.Printf("十六进制，小写字母，每字节两个字符 %X\n", d)

	fmt.Println("-------------------print bool with string --------------------")

	fmt.Printf("指针 %p\n", &d)


	fmt.Println("-------------------print bool with %%--------------------")
	fmt.Printf("%%\n", a)
	fmt.Printf("%%\n", b)
	fmt.Printf("%%\n", c)
	fmt.Printf("%%\n", d)
	fmt.Println("-------------------print bool with %%--------------------")
}
