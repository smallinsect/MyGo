// 22_if_elseif_else的使用
package main

import (
	"fmt"
)

func main() {
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else {
		fmt.Println("a != 10")
	}
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("a  10")
	}
}
