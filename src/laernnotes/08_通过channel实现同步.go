// 08_通过channel实现同步
package main

// //全局变量，创建一个channel
// var ch = make(chan int)

// //定一个打印机，参数为字符串，按每个字符打印
// func Printer(str string) {
// 	for _, data := range str {
// 		fmt.Printf("%c", data)
// 		time.Sleep(time.Second)
// 	}
// 	fmt.Printf("\n")
// }

// func person1() {
// 	Printer("hello")

// 	ch <- 666 //给管道写数据，发送
// }

// func person2() {
// 	<-ch //从管道取数据，接受，如果通道没有数据他就会阻塞
// 	Printer("world")
// }

// func main() {

// 	go person1()
// 	go person2()

// 	for {

// 	}
// }
