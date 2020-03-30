// 04_不定参数传递
package main

// func MyFunc01(args ...int) {
// 	for i, data := range args {
// 		fmt.Println(i, data)
// 	}
// }

// //不定参数
// func MyFunc02(args ...int) {
// 	//全部元素传递过去
// 	// MyFunc01(args...)
// 	//只想把后2个参数传递给另一个函数使用
// 	MyFunc01(args[1:3]...) //从args[1,3)所有的元素传递过去
// }

// func main() {
// 	MyFunc02(1, 2, 3, 4, 5)
// }
