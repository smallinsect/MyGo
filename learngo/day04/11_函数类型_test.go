// 11_函数类型
package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数取别名
type FuncType func(int, int) int //没有函数名，没有{}

func main() {
	var num int

	//声明一个函数类型，变量名叫fTest
	var fTest FuncType
	fTest = Add //是变量的赋值
	num = fTest(20, 10)
	fmt.Println(num)

	fTest = Sub
	num = fTest(20, 10)
	fmt.Println(num)
}
