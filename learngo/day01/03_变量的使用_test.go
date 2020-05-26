package main

import (
	"fmt"
	"testing"
)

func TestVar(t *testing.T) {
	//声明变量 var 变量名 类型，变量声明了，必须使用
	//
	var a int

	fmt.Println("a = ", a)

	//2、变量的初始化
	var b int = 10
	fmt.Println("b = ", b)
	b = 20
	fmt.Println("b = ", b)
	//3、自动推导类型，必须初始化，通过初始化的值确定的该类型
	c := 30
	fmt.Println("c = ", c)
	//%T打印变量所属的类型
	fmt.Printf("c type %T", c)
}
