// 13_字符串类型
package main

import (
	"fmt"
	"testing"
)

func TestStringType(t *testing.T) {
	var str1 string // 声明变量
	str1 = "abc"
	fmt.Println("str1 = ", str1)

	//自动推导类型
	str2 := "mike"
	fmt.Println("str2=", str2)
	fmt.Println(len(str2))
}
