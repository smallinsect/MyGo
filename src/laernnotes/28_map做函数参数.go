// 28_map做函数参数
package main

import (
	"fmt"
)

func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	fmt.Println("Hello World!")
	m1 := map[int]string{1: "mike", 2: "go", 3: "yoyo"}
	fmt.Println(m1)
	test(m1)
	fmt.Println(m1)
}
