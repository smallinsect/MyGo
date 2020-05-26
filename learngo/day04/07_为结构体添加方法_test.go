// 07_为结构体添加方法
package main

// type Person struct {
// 	name string //名字
// 	sex  byte   //性别
// 	age  int    //年龄
// }

// //带有接受者的函数叫方法
// func (tmp Person) PrintInfo() {
// 	fmt.Println(tmp.name)
// 	fmt.Println(tmp.sex)
// 	fmt.Println(tmp.age)
// }

// //通过一个函数，给成员赋值
// func (p *Person) SetInfo(n string, s byte, a int) {
// 	p.name = "aaaa"
// 	p.sex = 's'
// 	p.age = 22
// }
// func main() {
// 	fmt.Println("Hello World!")
// 	var p Person = Person{"fdsasd", 'd', 11}
// 	p.PrintInfo()

// 	var p1 Person
// 	(&p1).SetInfo("aaa", 'e', 22)
// 	p1.PrintInfo()
// }
