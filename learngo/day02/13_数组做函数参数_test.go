// 13_数组做函数参数
package main

import (
	"fmt"
)

//数组做参数，值拷贝
func modify(a [5]int) {
	a[0] = 666
	fmt.Println(a)
}

func main() {
	fmt.Println("Hello World!")

	a := [5]int{1, 2, 3, 4, 5} //初始化

	modify(a)
	fmt.Println(a)

}
