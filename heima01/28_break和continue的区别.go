// 28_break和continue的区别
package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for { //这个循环条件永远为真
		i++
		if i == 2 { //直接执行
			continue
		}
		if i > 5 { //跳出最近的一个循环
			break
		}
		fmt.Println(i)
		time.Sleep(time.Second) //延迟1秒

	}
}
