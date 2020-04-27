// 09_通过channel实现同步和数据交互
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	ch := make(chan string)

	go func() {
		defer fmt.Println("子协程结束")
		fmt.Println("1111111111")
		fmt.Println("2222222222")

		ch <- "我是子协程，工作完成"
	}()

	msg := <-ch
	fmt.Println(msg)
}
