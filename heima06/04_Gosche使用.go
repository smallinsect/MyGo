// 04_Gosche使用
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片断给协程执行时间
		runtime.Gosched()
		fmt.Println("hello")
	}
}
