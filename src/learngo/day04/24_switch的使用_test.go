// 24_switch的使用
package main

import (
	"fmt"
)

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println(num)
		// fallthrough//不跳出switch语句 相当于C语言不加break
	case 2, 3, 4:
		fmt.Println(num)
	default:
		fmt.Println(num)
	}

	switch num := 1; num {
	case 1:
		fmt.Println(num)
	case 2:
		fmt.Println(num)
	}

	switch { //可以没有条件
	case num > 10:
		fmt.Println(num)
	case num < 10:
		fmt.Println(num)
	case num == 10:
		fmt.Println(num)
	}
}
