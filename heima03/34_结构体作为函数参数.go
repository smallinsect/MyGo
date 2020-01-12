// 34_结构体作为函数参数
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

//值传递
func test(s Student) {
	s.id = 666
}

//地址传递
func test01(s *Student) {
	s.id = 666
}

func main() {
	fmt.Println("Hello World!")
	//顺序初始化，每个成员必须初始化
	var s1 Student = Student{1, "mki", 'm', 18, "bj"}
	fmt.Println(s1)
	test(s1)
	fmt.Println(s1)
	test01(&s1)
	fmt.Println(s1)
}
