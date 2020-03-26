// 11_方法的继承
package main

import (
	"fmt"
)

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

//Person类型，实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name=%s,sex=%c,age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承Person，字段成员和方法都继承
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

//覆盖父类函数
func (tmp *Student) PrintInfo() {
	fmt.Printf("%+v\n", tmp)
}

func main() {
	s := Student{Person{"hehe", 'm', 111}, 666, "bbbb"}
	s.PrintInfo()
	//显示调用
	s.Person.PrintInfo()
	fmt.Println("Hello World!")
}
