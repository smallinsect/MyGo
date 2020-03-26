// 05_结构体指针类型匿名字段
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
	int
	string
}

func main() {
	fmt.Println("Hello World!")
	s1 := &Student{Person{"mike", 'm', 18}, 1, "bbbb"}
	fmt.Println(s1)
	//%+v 显示更详细
	fmt.Printf("%+v\n", s1)
	//指针匿名类成员访问
	s1.name = "aaaa"
	//指针匿名字段操作
	s1.int = 22
	s1.string = "dddd"

	//指定初始化成员，没有初始化的成员
	s4 := Student{Person: Person{name: "mike"}, id: 1}
	fmt.Println(s4)
}
