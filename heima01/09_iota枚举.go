// 09_iota枚举.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	// iota常量自动生成器，每个一行，自动累加1
	//iota给常量赋值使用
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a=%d,b=%d,c=%d\n", a, b, c)

	// iota遇到const，重置为0
	const d = iota
	fmt.Printf("d=%d\n", d)

	// 可以只写一个iota
	const (
		e = iota
		f
		g
	)
	fmt.Printf("e=%d,f=%d,g=%d\n", e, f, g)

	// 如果是同一行，值都一样
	const (
		h       = iota
		j, k, l = iota, iota, iota
		m       = iota
	)
	fmt.Printf("h=%d,j=%d,k=%d,l=%d,m=%d\n", h, j, k, l, m)
}
