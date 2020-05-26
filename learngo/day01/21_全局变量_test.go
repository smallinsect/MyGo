// 21)全局变量
package main

import (
	"fmt"
	"testing"
)

//定义一个全局变量
var a int

func TestQuanju(t *testing.T) {
	fmt.Println(a)
}
