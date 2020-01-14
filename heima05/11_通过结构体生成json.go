// 11_通过结构体生成json
package main

import (
	"encoding/json"
	"fmt"
)

/*

{
	"company":"itcast",
	"subjects":[
		"Go",
		"C++",
		"Python",
		"Test"
	],
	"isok":true,
	"price": 666.666
}

*/
//成员变量名首字符必须大写
type IT struct {
	Company  string
	Subjects []string
	Isok     bool
	Price    float64
}

//二次编码
type IT1 struct {
	Company  string   `json:"-"` //此字段不显示到屏幕上
	Subjects []string `json:"subjects"`
	Isok     bool     `json:",string"` //以字符串形式显示
	Price    float64  `json:"price"`
}

func main() {
	fmt.Println("Hello World!")
	//定义一个结构体变量，同事初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}

	//编码，根据内容生成json文本
	buf, err := json.Marshal(s)
	buf1, err1 := json.MarshalIndent(s, "", " ") //格式化编码

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(buf))
	}

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(string(buf1))
	}

	s1 := IT1{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
	buf2, err2 := json.MarshalIndent(s1, "", " ") //格式化编码
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(string(buf2))
	}
}
