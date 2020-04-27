// 10_无缓存的channel
package main

import (
	"fmt"
	"time"
)

func main() {

	//创建一个无缓存的通道
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("i=", i)
			//往管道写入数据
			ch <- i
		}
	}()
	//主协程沉睡2秒
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		//读取管道中的数据
		num := <-ch
		fmt.Println("num=", num)
	}
}
