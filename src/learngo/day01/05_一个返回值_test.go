// 04_不定参数传递
package main

import (
	"fmt"
	"testing"
)

//无参返回值
func ReturnFunc01() int {
	return 666
}

//给返回值起一个变量名，go推荐写法
func ReturnFunc02() (result int) {
	result = 666
	return
}

func TestReturn(t *testing.T) {
	var a int
	a = ReturnFunc01()

	b := ReturnFunc02()

	fmt.Println(a)
	fmt.Println(b)
}
