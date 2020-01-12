// 18_defer和匿名函数结合使用
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	a := 1
	b := 2
	//先把参数传入进去
	defer func(a, b int) {
		fmt.Println(a, b)
	}(a, b)

	a = 11
	b = 22
	fmt.Println(a, b)
}
