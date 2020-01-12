// 15_闭包的特点
package main

import (
	"fmt"
)

func test01() int {
	var x int
	x++
	return x * x
}

//函数的返回值是一个匿名函数，返回一个函数类型
func test02() func() int {
	var x int //没有初始化，值为0
	//形成一个闭包
	return func() int {
		x++
		return x * x
	}
}

func main() {
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())

	//通过f调用闭包
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
