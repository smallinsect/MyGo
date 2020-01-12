// 22_切片做函数参数
package main

import (
	"fmt"
)

func InitData(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
}

func main() {
	fmt.Println("Hello World!")
	s := make([]int, 5)
	InitData(s)
	fmt.Println(s)
}
