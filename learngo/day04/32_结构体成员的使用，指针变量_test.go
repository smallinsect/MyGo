// 32_结构体成员的使用，指针变量
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
	var p1 *Student
	fmt.Println(s1)
	p1 = &s1
	p1.id = 222
	p1.name = "small"
	fmt.Println(s1)

	//通过new申请一块内存
	p2 := new(Student)
	fmt.Println(*p2)
}
