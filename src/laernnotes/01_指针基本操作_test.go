// 01_指针基本操作
package main

import (
	"fmt"
	"testing"
)

// 测试指针
func TestPointer(t *testing.T) {
	fmt.Println("Hello World!")
	//每一个变量有2层含义：变量的内存，和变量的地址
	var a int
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %v\n", &a)

	//定一个指针变量
	var p *int
	p = &a
	fmt.Printf("p = %v\n", p)
	fmt.Printf("&p = %v\n", &p)
	fmt.Printf("*p = %d\n", *p)
}
