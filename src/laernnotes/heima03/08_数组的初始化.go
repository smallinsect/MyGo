// 08_数组的初始化
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//1、全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := [5]int{6, 7, 8, 9, 10}
	fmt.Println(b)

	//部分初始化
	c := [5]int{1, 2, 3}
	fmt.Println(c)

	//指定某个元素初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Println(d)

}
