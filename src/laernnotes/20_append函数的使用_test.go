// 20_append函数的使用
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	s1 := []int{}
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1)

	s1 = append(s1, 5)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1)
}
