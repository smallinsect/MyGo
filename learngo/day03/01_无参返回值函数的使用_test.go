// 01_无参返回值函数的使用
package main

import (
	"fmt"
	"testing"
)

func MyFunc() {
	aaa := 666
	fmt.Println(aaa)
}

// 测试函数
func TestFunc(t *testing.T) {
	fmt.Println("Hello World!")
	//函数调用
	MyFunc()
}
