// 27_map删除
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	m1 := map[int]string{1: "mike", 2: "go", 3: "yoyo"}
	fmt.Println(m1)

	delete(m1, 1) //删除键值为1的内容
	fmt.Println(m1)
}
