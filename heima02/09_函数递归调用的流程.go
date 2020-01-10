// 09_函数递归调用的流程
package main

import (
	"fmt"
)

func funca(i int) (a int) {
	if i > 5 {
		a = 1
		return
	}
	a = funca(i + 1)
	fmt.Println(i + 1)
	return
}

func main() {
	funca(0)
}
