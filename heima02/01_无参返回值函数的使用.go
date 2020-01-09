// 01_无参返回值函数的使用
package main

import (
	"fmt"
)

func MyFunc() {
	aaa := 666
	fmt.Println(aaa)
}

func main() {
	fmt.Println("Hello World!")
	//函数调用
	MyFunc()
}
