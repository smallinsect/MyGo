// 08_普通函数的调用流程
package main

import (
	"fmt"
)

func funcb() (a int) {
	a = 10
	return
}

func funca() (a int) {
	a := funcb()
	return
}

func main() {
	funca()
}
