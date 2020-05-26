// 04_不定参数传递
package main

import (
	"fmt"
	"testing"
)

func Func01(args ...int) {
	for i, data := range args {
		fmt.Println(i, data)
	}
}

//不定参数
func Func02(args ...int) {
	//全部元素传递过去
	// MyFunc01(args...)
	//只想把后2个参数传递给另一个函数使用
	Func01(args[1:3]...) //从args[1,3)所有的元素传递过去
}

func TestTooVar(t *testing.T) {
	Func02(1, 2, 3, 4, 5)
}
