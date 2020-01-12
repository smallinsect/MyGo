// 10_数组比较和赋值
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//支持比较，支持==和!=，比较是不是每一个元素都一样，2个数组比较，和数组类型要一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a == c)
	fmt.Println(a != c)

	//同类型的数组可以赋值
	var d [5]int
	d = a
	fmt.Println(d)
}
