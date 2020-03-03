// 03_同名字段
package main

import (
	"fmt"
)

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
	name   string
}

func main() {
	fmt.Println("Hello World!")
	var s1 Student
	s1.name = "small"
	fmt.Printf("%+v\n", s1)

	//显示调用
	s1.Person.name = "insect"
	fmt.Printf("%+v\n", s1)
}
