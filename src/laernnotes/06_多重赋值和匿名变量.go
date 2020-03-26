package main

import "fmt"

//go函数可以返回多个值
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	//多重赋值
	a, b, c := 10, 20, 30
	fmt.Println(a, b, c)

	i := 10
	j := 20
	//交换两个变量的值
	i, j = j, i
	fmt.Println(i, j)

	var tmp int
	//匿名变量，丢弃数据不作处理,_匿名变量配合函数返回值使用，才有优势
	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

	a, b, c = test()
	fmt.Println(a, b, c)
	_, b, _ = test()
	fmt.Println(a, b, c)
}
