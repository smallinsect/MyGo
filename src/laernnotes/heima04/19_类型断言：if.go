// 19_类型断言：if
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
		//第一个返回时值，第二个返回判断结构的真假
		if value, ok := data.(int); ok == true {
			fmt.Println(idx, value, data, "整形")
		} else if value, ok := data.(string); ok == true {
			fmt.Println(idx, value, data, "字符串")
		} else if value, ok := data.(Student); ok == true {
			fmt.Println(idx, value, data, "Student")
		}
	}
}
