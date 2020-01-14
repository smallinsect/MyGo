// 14_json解析到map
package main

import (
	"encoding/json"
	"fmt"
)

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
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonbuf), &m) //第二个参数要地址传递
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}
	var str string
	//类型断言
	for key, value := range m {
		fmt.Println(key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Println("string=", value, str)
		case bool:
		}
	}
}
