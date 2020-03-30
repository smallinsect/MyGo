// 02_普通参数列表
package main

import (
	"fmt"
	"testing"
)

func MyFunc11(a int) {
	fmt.Println(a)
	a = 111
	fmt.Println(a)
}

func MyFunc01(a int, b int) {
	fmt.Println(a, b)
}

func MyFunc02(a int, b float32, c string) {
	fmt.Println(a, b, c)
}

//a b c都是string
func MyFunc03(a, b, c string) {
	fmt.Println(a, b, c)
}

// 测试普通函数
func TestFunc01(t *testing.T) {
	fmt.Println("Hello World!")
	MyFunc11(333)
	MyFunc01(11, 22)
}
