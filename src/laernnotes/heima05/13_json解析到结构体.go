// 13_json解析到结构体
package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"` //此字段不显示到屏幕上
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"` //以字符串形式显示
	Price    float64  `json:"price"`
}

func main() {
	fmt.Println("Hello World!")
	jsonbuf := `
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
	`
	var tmp IT                                   //定一个结构体变量
	err := json.Unmarshal([]byte(jsonbuf), &tmp) //第二个参数要地址传递
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tmp)
	}
}
