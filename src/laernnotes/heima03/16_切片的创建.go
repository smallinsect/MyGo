// 16_切片的创建
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//切片和数组的区别
	// a := [5]int{}

	//切片 长度和容量不固定
	s := []int{}
	s = append(s, 11)
	fmt.Println(s)

	//借助make函数 make(类型 长度 容量)
	s2 := make([]int, 5, 10)
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 5)
	fmt.Println(s3, len(s3), cap(s3))
}
