// 27_range的使用
package main

import (
	"fmt"
)

func main() {

	str := "abc"
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	//迭代打印每个元素，默认返回2个值：一个是元素位置，一个是元素本身
	for i, data := range str {
		fmt.Println(i, data)
	}
	for i := range str { //只使用下标，丢弃元素
		fmt.Println(i, data)
	}
	for i, _ := range str {
		fmt.Println(i, data)
	}
}
