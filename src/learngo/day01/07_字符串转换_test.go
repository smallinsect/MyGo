// 07_字符串转换
package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestStringChange(t *testing.T) {
	fmt.Println("Hello World!")
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数，第3个为指定10进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abced")

	// fmt.Println(strings(slice)) //转换为字符串后打印

	//其它类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	//'f' 指打印格式，以小数点方式， -1指小数点位数（紧缩模式）， 64以float64处理
	// str = strconv.AppendFloat(3.14, 'f', -1, 64)
	// fmt.Println(str)

	//整型转字符串，常用
	str = strconv.Itoa(666)
	fmt.Println(str)

	//字符串转其他类型
	flag, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(flag)
	}

	//把字符串转换为整型
	a, _ := strconv.Atoi("22222")
	fmt.Println(a)
}
