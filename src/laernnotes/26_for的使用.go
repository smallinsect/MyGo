// 26_for的使用
package main

import (
	"fmt"
)

func main() {
	// for 初始化条件; 判断条件; 条件变化
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
