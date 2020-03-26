// 08_值语义和引用语义
package main

import (
	"fmt"
)

type Person struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

//值语义
func (p Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age

	fmt.Printf("%p\n", &p.age)
}

//接受者为指针变量，引用传递
func (p *Person) SetInfoPointer(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age

	fmt.Printf("%p\n", &p.age)
}

func main() {
	fmt.Println("Hello World!")
	var p1 Person

	p1.SetInfo("hehheh", 'd', 222)
	fmt.Println(p1)

	(&p1).SetInfoPointer("hehheh", 'd', 222)
	fmt.Println(p1)
}
