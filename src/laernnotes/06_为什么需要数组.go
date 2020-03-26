// 06_为什么需要数组
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	// id1 := 1
	// id2 := 2
	// id3 := 3

	//定义一个数组
	var id [50]int

	//操作数组，通过下标，从0开始，到len()-1
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Println(id[i])
	}
}
