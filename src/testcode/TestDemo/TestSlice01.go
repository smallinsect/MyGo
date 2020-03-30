package testdemo

import (
	"fmt"
)

// SliceFunc ...
func SliceFunc() {
	// slice := make([]int, 0)

	// for i := 0; i < 100000000; i++ {
	// 	slice = append(slice, i)
	// }
	// fmt.Println(slice)
	s := make([]int, 0)

	for i := 11; i < 20; i++ {
		s = append(s, i)
	}
	s = append(s, 111)
	fmt.Println(s)
	// idx := 0
	// s = append(s[:idx], s[idx+1:]...)
	fmt.Println(s[:1])
}

// // MapFunc01 ... map的使用
// func MapFunc01() {
// 	//map需要使用make分配空间，才能使用
// 	var m1 map[string]string
// 	m1 = make(map[string]string, 10)
// 	m1["11111"] = "小昆虫"
// 	m1["22222"] = "小白菜"
// 	m1["33333"] = "小东风"
// 	fmt.Println("m1=", m1)

// 	//类型推导
// 	m2 := make(map[string]string, 10)
// 	m2["11111"] = "小昆虫"
// 	m2["22222"] = "小白菜"
// 	m2["33333"] = "小东风"
// 	fmt.Println("m2=", m2)

// 	//类型推导，初始化
// 	m3 := map[string]string{
// 		"11111": "小昆虫",
// 		"22222": "小白菜",
// 		"33333": "小东风",
// 	}
// 	fmt.Println("m3=", m3)
// }

// // MapFunc02 ... map的增删改查
// func MapFunc02() {
// 	m := make(map[string]string)
// 	m["11111"] = "小昆虫"
// 	m["22222"] = "小白菜"
// 	m["33333"] = "小东风"
// 	fmt.Println("m=", m)
// 	//如果map中没有键值"44444"，就新增
// 	m["44444"] = "小青蛙"
// 	fmt.Println("m=", m)
// 	//如果map中有键值"11111"，就修改
// 	m["11111"] = "小昆虫*"
// 	fmt.Println("m=", m)
// 	//删除map中的键值"22222"的数据
// 	delete(m, "22222")
// 	fmt.Println("m=", m)
// 	//删除map中所有的数据
// 	//1、遍历所有key，逐个删除
// 	//2、直接重新make新空间
// 	// m = make(map[string]string)
// 	//查找键值["11111"
// 	val, ok := m["11111"]
// 	if ok {
// 		fmt.Println("11111", val)
// 	} else {
// 		fmt.Println("没有11111")
// 	}
// }

// // MapFunc03 ... map的range遍历
// func MapFunc03() {
// 	m := make(map[string]string)
// 	m["11111"] = "小昆虫"
// 	m["22222"] = "小白菜"
// 	m["33333"] = "小东风"
// 	m["44444"] = "小青蛙"
// 	fmt.Println("m=", m)
// 	//range遍历map
// 	for k, v := range m {
// 		fmt.Println("k=", k, "v=", v)
// 	}

// 	//复杂结构的map
// 	m1 := make(map[string]map[string]string)
// 	m1["aaaaa"] = make(map[string]string)
// 	m1["bbbbb"] = make(map[string]string)
// 	m1["aaaaa"]["11111"] = "小昆虫"
// 	m1["aaaaa"]["22222"] = "小白菜"
// 	m1["bbbbb"]["33333"] = "小东风"
// 	m1["bbbbb"]["44444"] = "小青蛙"
// 	fmt.Println("m1=", m1)
// 	for k1, v1 := range m1 {
// 		fmt.Println(k1, v1)
// 		for k2, v2 := range v1 {
// 			fmt.Println("", k2, v2)
// 		}
// 	}
// }
