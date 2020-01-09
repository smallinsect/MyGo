// 18_类型转换
package main

import (
	"fmt"
)

func main() {
	var flag bool
	flag = true

	fmt.Println(flag)
	fmt.Printf("%d\n", flag)
	//bool类型不能转换为整形
	// fmt.Printf("%d\n", int(flag))
	//整形不能转换为bool类型
	// 这种转换的类型，叫不兼容类型

	var ch byte
	ch = 'a'
	var t int
	t = int(ch)
	fmt.Println(t)
}
