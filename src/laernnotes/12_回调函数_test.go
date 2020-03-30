// 12_回调函数
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

func Calc(a, b int, fTest FuncType) (result int) {
	result = fTest(a, b) //是变量的赋值
	return
}

func main() {
	fmt.Println(Calc(222, 111, Add))
	fmt.Println(Calc(222, 111, Sub))
}
