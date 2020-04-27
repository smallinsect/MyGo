package test

import (
	"fmt"
)

//结构体 首字母小写，只能在同一个报名中使用
type student struct {
	id   int
	name string
}

//结构体 首字母大写，可以挎包使用
type Student struct {
	Id   int
	Name string
}

//函数名首写字母小写，只能在同一个包名使用
func myFunc() {
	fmt.Println("myFunc()")
}

//函数名首写字母大写，可以别的包名使用
func MyFunc() {
	fmt.Println("MyFunc()")
}
