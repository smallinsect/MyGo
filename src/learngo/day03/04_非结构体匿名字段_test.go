// 04_非结构体匿名字段
package main

// type mystr string //自定义类型，给类型改名字

// type Person struct {
// 	name string //名字
// 	sex  byte   //性别
// 	age  int    //年龄
// }

// type Student struct {
// 	Person //只有类型，没有名字，匿名字段，继承了Person的成员
// 	int
// 	string
// }

// func main() {
// 	fmt.Println("Hello World!")
// 	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bbbb"}
// 	fmt.Println(s1)
// 	//%+v 显示更详细
// 	fmt.Printf("%+v\n", s1)
// 	s1.name = "heheh"
// 	//匿名成员操作
// 	s1.int = 10
// 	s1.string = "aaa"

// 	//指定初始化成员，没有初始化的成员
// 	s4 := Student{Person: Person{name: "mike"}, id: 1}
// 	fmt.Println(s4)
// }
