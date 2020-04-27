// 15_接口的定义和实现
package main

import (
	"fmt"
)

//定义接口类型
type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

//Student实现了此方法
func (tmp *Student) sayhi() {
	fmt.Println(tmp)
}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Println(tmp)
}

func main() {
	fmt.Println("Hello World!")
	//定义接口变量
	var h Humaner
	h = &Student{"hhh", 11}
	h.sayhi()

	h = &Teacher{"aaaa", "bbb"}
	h.sayhi()
}
