// 20_局部变量
package main

import (
	"fmt"
	"testing"
)

func TestNeibuVar(t *testing.T) {
	fmt.Println("Hello World!")
	//定义在大括号里面的都是局部变量
	//作用域，变量其作用的范围
	a := 111
	{
		i := 2222
		fmt.Println(i)
	}
	fmt.Println(a)

	if flag := 3; flag > 0 {
		fmt.Println(flag)
	}
}
