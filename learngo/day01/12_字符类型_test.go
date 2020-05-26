// 12_字符类型
package main

import (
	"fmt"
	"testing"
)

func TestStringVar(t *testing.T) {
	var ch byte //声明字符类型
	ch = 97
	//格式化输出，%c %d
	fmt.Printf("%c %d\n", ch, ch)
}
