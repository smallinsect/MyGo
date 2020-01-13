// 06_面向过程和对象函数的区别
package main

import (
	"fmt"
)

//面向过程
func Add01(a, b int) int {
	return a + b
}

//面向对象，方法：给某个类型绑定一个函数
type long int

//tmp叫接受者，接受者就是传递的一个参数
func (tmp long) Add02(other long) long {
	return tmp + other
}

func main() {
	fmt.Println("Hello World!")
	var result1 int
	result1 = Add01(11, 22)
	fmt.Println(result1)

	//定义一个变量
	var a long = 1
	result2 := a.Add02(2)
	fmt.Println(result2)
}
