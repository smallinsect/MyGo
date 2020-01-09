// 16_格式化输出
package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.124

	fmt.Printf("%T %T %T %T\n", a, b, c, d)
	//%d 整形格式
	//%c 字符格式
	//%s 字符格式
	//%v 自动类型
}
