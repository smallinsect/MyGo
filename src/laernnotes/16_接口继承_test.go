// 16_接口继承
package main

import (
	"fmt"
)

type Humaner interface {
	sayhi()
}

//继承接口
type Man interface {
	Humaner
	wall()
}

type Student struct {
	name string
}

//实现接口
func (tmp *Student) sayhi() {
	fmt.Println("sayhi Student")
}

func (tmp *Student) wall() {
	fmt.Println("wall Student")
}

type Teacher struct {
	name string
}

//实现接口
func (tmp *Teacher) sayhi() {
	fmt.Println("sayhi Teacher")
}

func (tmp *Teacher) wall() {
	fmt.Println("wall Teacher")
}

func main() {
	//超集可以转换为子集，反过来，不可以。
	var man Man
	man = &Student{"我是学生"}
	man.sayhi()
	man.wall()

	man = &Teacher{"我是老师"}
	man.sayhi()
	man.wall()
}
