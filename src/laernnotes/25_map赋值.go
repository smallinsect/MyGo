// 25_map赋值
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	m1 := map[int]string{1: "aaa", 2: "bbb"}
	fmt.Println(m1)

	m1[1] = "c++"
	fmt.Println(m1)

	m1[3] = "go"
	fmt.Println(m1)
}
