// 16_defer的使用
package main

import (
	"fmt"
)

func main() {
	//延迟调用，main结束前调用，类似析构函数
	defer fmt.Println("Hello2")

	fmt.Println("Hello1")
}
