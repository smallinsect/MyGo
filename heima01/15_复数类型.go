// 15_附属类型
package main

import (
	"fmt"
)

func main() {
	//声明变量
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println(t)

	//自动推导类型
	t1 := 3.3 + 4.4i
	fmt.Println(t1)

	//实部 虚部
	fmt.Println(real(t1), imag(t1))
}
