// 29_结构体普通变量初始化
package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	fmt.Println("Hello World!")
	//顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "mki", 'm', 18, "bj"}
	fmt.Println(s1)

	//指定成员初始化
	s2 := Student{name: "small", addr: "bb"}
	fmt.Println(s2)
}
