// 25_recover的使用
package main

import (
	"fmt"
)

func test01() {
	fmt.Println("aaaaaaaaa")
}
func test02(idx int) {
	fmt.Println("bbbbbbbbb")
	defer func() { //定义延迟函数
		//恢复异常
		if err := recover(); err != nil {
			//有数组越界，就打印信息
			fmt.Println(err)
		}
	}()
	var arr [10]int
	arr[idx] = 100
}
func test03() {
	fmt.Println("ccccccccc")
}

func main() {
	fmt.Println("Hello World!")
	test01()
	test02(111)
	test03()
}
