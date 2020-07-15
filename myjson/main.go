package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// omitempty 该变量为空串 为0 为false 不进行转换
type Test struct {
	Name string `json:"name,omitempty"`
	Sex  bool   `json:"sex,omitempty"`
}

func main1() {
	// test1 := Test{
	// 	Name: "111",
	// 	Sex:  true,
	// }
	// byte1, err := json.Marshal(test1)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(byte1))
	// }
	var v Test
	j := `{"name":"","sex":true}`
	err := json.Unmarshal([]byte(j), &v)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v)
	}
}

type User struct {
	Name       string  `json:"name"`
	Age        int64   `json:"age"`
	Account_id string  `json:"account_id"`
	Password   string  `json:"password"`
	RMB        float64 `json:"RMB"`
	Sex        bool    `json:"sex"`
}

func main() {
	bytes, err := ioutil.ReadFile("./MyUsers.json")
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}
	u := &User{}
	err = json.Unmarshal(bytes, u)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return
	}
	fmt.Printf("%+v\n", u)
}
