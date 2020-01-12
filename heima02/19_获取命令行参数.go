// 19_获取命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	//接受用户传递的参数
	list := os.Args
	n := len(list)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		fmt.Println(list[i])
	}
	for idx, data := range list {
		fmt.Println(idx, data)
	}
	fmt.Println("Hello World!")
}
