package main

import "fmt"

func main() {
	var a int
	//赋值前，必须定义
	a = 10
	fmt.Println("a = ", a)
	//:= 自动推导类型，先声明变量b，再给b赋值为20
	b := 20
	fmt.Println("b = ", b)
}
