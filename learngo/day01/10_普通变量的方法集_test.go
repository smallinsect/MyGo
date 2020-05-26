// 10_普通变量的方法集
package main

import (
	"fmt"
	"testing"
)

// type Person struct {
// 	name string //名字
// 	sex  byte   //性别
// 	age  int    //年龄
// }

// //值语义
// func (p Person) SetInfo() {
// 	fmt.Println("SetInfo")
// }

// //接受者为指针变量，引用传递
// func (p *Person) SetInfoPointer() {
// 	fmt.Println("SetInfoPointer")
// 	p.name = "1111"
// }

func TestPutongVar(t *testing.T) {
	fmt.Println("Hello World!")
	p := &Person{"mmm", 11, 222}

	//直接调用
	p.SetInfo()
	//先转换成指针，再调用
	p.SetInfoPointer()
	fmt.Println(p)
}
