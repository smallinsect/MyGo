// 26_map遍历
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	m1 := map[int]string{1: "mike", 2: "go", 3: "yoyo"}

	for key, value := range m1 {
		fmt.Println(key, value)
	}
	//判断键值对是否存在
	value, ok := m1[4]
	if ok == true {
		fmt.Println("m1[4] 存在", value)
	} else {
		fmt.Println("m1[4] 不存在")
	}

}
