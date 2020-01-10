// 13_匿名函数和闭包
package main

import (
	"fmt"
)

func main() {
	a := 111
	b := "abc"

	//匿名函数，没有函数名字
	f1 := func() { //形成一个闭包
		fmt.Println(a)
		fmt.Println(b)
	}

	//调用f1函数
	f1()

	//给一个函数类型起别名
	type FuncType func()
	//声明变量
	var f2 = FuncType
	//变量赋值
	f2 = f1()
	//执行函数
	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Println("小昆虫")
	}()

	f3 := func(i, j int) {
		fmt.Println(i, j)
	}
	//调用
	f3()

	a, b := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(11, 22) //()代表直接调用
}
