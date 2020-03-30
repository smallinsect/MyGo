// 11_浮点型
package main

import (
	"fmt"
)

func main() {
	//1、声明变量
	var f1 float32 = 3.124
	fmt.Println("f1=", f1)

	//推导类型默认float64
	f2 := 3.14
	fmt.Printf("f2 type %T\n", f2)
	fmt.Println("Hello World!")
}
