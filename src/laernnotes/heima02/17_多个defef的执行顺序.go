// 17_多个defef的执行顺序
package main

import (
	"fmt"
)

func test(x int) {
	result := 100 / x
	fmt.Println(result)
}

func main() {
	defer fmt.Println("aaaaaaa")
	defer fmt.Println("bbbbbbb")

	defer test(0)

	defer fmt.Println("fffffff")
}
