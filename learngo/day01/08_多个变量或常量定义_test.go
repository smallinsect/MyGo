// 08_多个变量或常量定义.go
package main

import (
	"fmt"
	"testing"
)

func TestTooVar01(t *testing.T) {
	fmt.Println("Hello World!")
	// 不同类型变量的声明（定义）
	// var a int
	// var b float64
	var (
		a int
		b float64
	)
	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	const (
		i = 3.14
		j = 6.15
	)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("j = %T\n", j)
}
