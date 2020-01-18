// 06_GOMAXPROCS的使用
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World!")
	n := runtime.GOMAXPROCS(1)
	for {
		go fmt.Println(1)
		fmt.Println(0)
	}
}
