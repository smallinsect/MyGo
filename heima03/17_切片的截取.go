// 17_切片的截取
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max] 取消表low开始，len=high-low，容量=max-low

	s1 := array[:]
	//...看帮助文档
}
