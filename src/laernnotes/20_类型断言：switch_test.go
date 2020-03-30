// 20_类型断言：switch
package main

import (
	"fmt"
)

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                  //int
	i[1] = "hello go"         //string
	i[2] = Student{"mmm", 66} //Studen
	//类型查询，类型断言
	for idx, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Println(idx, value, data, "整形")
		case string:
			fmt.Println(idx, value, data, "字符串")
		case Student:
			fmt.Println(idx, value, data, "Student")
		}
	}
}
