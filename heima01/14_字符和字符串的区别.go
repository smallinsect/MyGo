// 14_字符和字符串的区别
package main

import (
	"fmt"
)

func main() {
	var ch byte
	var str string

	ch = 'a'
	str = "a"
	fmt.Println(ch)
	fmt.Println(str)

	str = "asdfghjkkl"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}
}
