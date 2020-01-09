// 21_if的使用
package main

import (
	"fmt"
)

func main() {
	s := "王思聪"
	if s == "王思聪" {
		fmt.Println("左右手各一个妹子")
	}
	//if支持1个初始化，初始化和判断条件以分割符隔开
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}
}
