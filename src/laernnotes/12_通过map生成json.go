// 12_通过map生成json
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//创建一个map
	m := make(map[string]interface{}, 4)
	m["Company"] = "itcast"
	m["Subjects"] = []string{"Go", "C++", "Python", "Test"}
	m["Isok"] = true
	m["Price"] = 666.666

	//编码成json
	result, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(result))
	}
}
