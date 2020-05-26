// 09_二维数组的介绍
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//有多少个方括号，就有多少维
	var a [3][4]int

	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
		}
	}
	fmt.Println(a)

	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(b)

	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(c)

	d := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println(d)

	e := [3][4]int{0: {1, 2, 3, 4}, 2: {5, 6, 7, 8}}
	fmt.Println(e)
}
