// 06_字符串操作
package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("Hello World!")
	//是否包含
	fmt.Println(strings.Contains("Hello World!", "Hello"))
	//Joins组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "x")
	fmt.Println(buf)

	//Index 查找子串的位置
	fmt.Println(strings.Index("Hello World!", "Hello"))
	fmt.Println(strings.Index("Hello World!", "GO"))

	buf = strings.Repeat("go", 3)
	fmt.Println(buf) //gogogo

	//Split 以指定的分隔符拆分
	buf = "hello@world"
	s2 := strings.Split(buf, "@")
	fmt.Println(s2)

	//Trim去掉两头的字符
	buf = strings.Trim("   ar e ok ?  ", " ") //去掉首尾空格
	fmt.Println(buf)

	buf1 := strings.Fields("   ar e ok ?  ")
	fmt.Println(buf1)
}
