// 09_指针变量的方法集
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
func (p Person) SetInfo() {
	fmt.Println("SetInfo")
}

//接受者为指针变量，引用传递
func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	fmt.Println("Hello World!")
	p := &Person{"mmm", 11, 222}

	//内部做的转换，先把指针p，转成*p后再调用
	p.SetInfo()
	//指针直接调用
	p.SetInfoPointer()
}
