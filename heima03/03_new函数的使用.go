// 03_new函数的使用
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	var p *int
	//内存只管申请，不管释放，go语言有内存释放内存
	p = new(int)
	*p = 666
	fmt.Println(*p)
	//自动推导类型
	q := new(int)
	*q = 777
	fmt.Println(*q)
}
