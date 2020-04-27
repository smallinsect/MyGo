// 14_数组指针做函数参数
package main

import (
	"fmt"
)

//数组指针，指向数组
func modify(p *[5]int) {
	(*p)[0] = 666
	fmt.Println(*p)
}

func main() {
	fmt.Println("Hello World!")

	a := [5]int{1, 2, 3, 4, 5} //初始化

	modify(&a)
	fmt.Println(a)

}
