// 24_map的基本使用
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//定一个变量，类型为map[int]string
	var m1 map[int]string
	fmt.Println(m1)
	fmt.Println(len(m1))

	m2 := make(map[int]string)
	fmt.Println(m2)
	fmt.Println(len(m2))

	m3 := make(map[int]string, 5)
	m3[1] = "mike"
	m3[2] = "go"
	m3[3] = "c++"
	fmt.Println(m3)
	fmt.Println(len(m3))

	m4 := map[int]string{1: "dd", 2: "222"}
	fmt.Println(m4)
	fmt.Println(len(m4))

}
