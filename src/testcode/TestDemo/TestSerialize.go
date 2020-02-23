package testdemo

import (
	"encoding/json"
	"fmt"
)

// Monster ...
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

// SerializeFunc ... 序列化
func SerializeFunc() {
	monster := Monster{
		Name:     "小昆虫",
		Age:      10000,
		Birthday: "0000-00-00",
		Sal:      22.33,
		Skill:    "飞沙走石",
	}
	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	}
	//输出序列化结果
	fmt.Println("序列化之后，data=", data)
	fmt.Println("序列化之后，data=", string(data))
}

// SerializeFunc01 ... map序列化
func SerializeFunc01() {
	//使用map，需要make
	m := make(map[string]interface{})
	m["name"] = "小昆虫"
	m["age"] = 60000
	m["address"] = "星辰大海"
	//将m序列化
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	}
	//输出序列化结果
	fmt.Println("序列化之后，data=", string(data))
}

// SerializeFunc02... []map序列化
func SerializeFunc02() {
	var sm []map[string]interface{}
	//使用map，需要make
	m1 := make(map[string]interface{})
	m1["name"] = "小昆虫"
	m1["age"] = 666666
	m1["address"] = "星辰大海"

	m2 := make(map[string]interface{})
	m2["name"] = "小白菜"
	m2["age"] = 888888
	m2["address"] = "流浪银河"

	//将m1和m2放到切片中
	sm = append(sm, m1)
	sm = append(sm, m2)

	//将切片序列化
	data, err := json.Marshal(sm)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	}
	//输出序列化结果
	fmt.Println("序列化之后，data=", string(data))
}

// SerializeFunc03 ... 基本数据类型序列化
func SerializeFunc03() {
	var f float64 = 66.66
	data, err := json.Marshal(f)
	if err != nil {
		fmt.Println("序列化失败，err=", err)
		return
	}
	//输出序列化结果
	fmt.Println("序列化之后，data=", string(data))
}
