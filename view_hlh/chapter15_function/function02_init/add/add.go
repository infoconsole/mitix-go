package add

// _ 表示只是导入了这个包做了初始化，但是没有引用包里的任何东西
//导入的包最先进行初始化
import (
	"fmt"
	_ "mitix-go/view_hlh/chapter15_function/function02_init/test"
)

var Name = "Hello"
//小写变量无法访问
var age = 22

const Names string = "const info"

//const Names="const info" 简写

/**
1.全局变量的声明先被执行
2.每个源文件都可以包含一个init函数
类似于java里面的静态代码块
init方法会在初始化的时候被执行 在全局变量执行以后
*/
func init() {
	fmt.Println("init  Name is change to Hello init")
	Name = "Name Change init"
	age = 88
}

func TestInit() {
	Name = "Name Change TestInit"
	age = 99
}
