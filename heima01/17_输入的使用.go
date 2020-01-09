// 17_输入的使用
package main

import (
	"fmt"
)

func main() {
	var a int //声明一个变量
	fmt.Println("请输入一个整形")
	fmt.Scanf("%d", &a)
	fmt.Println("a=", a)

	fmt.Scan(&a)
	fmt.Println("a=", a)
}
