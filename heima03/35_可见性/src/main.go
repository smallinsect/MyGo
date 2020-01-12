package main

import (
	"fmt"
	"test"
)

func main() {
	// test.myFunc()
	test.MyFunc()
	fmt.Println("aaaaaaaaaaa")

	// var s test.student
	var s test.Student
	s.Id = 10
	s.Name = "fdasd"
	fmt.Println(s)
}
