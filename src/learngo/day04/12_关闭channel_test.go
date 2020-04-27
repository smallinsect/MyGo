// 12_关闭channel
package main

import (
	"fmt"
	"time"
)

func main() {

	//创建一个无缓存的通道
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("i=", i)
			//往管道写入数据
			ch <- i
		}
		//关闭管道
		close(ch)
	}()
	//主协程沉睡2秒
	time.Sleep(2 * time.Second)
	// for {
	// 	//读取管道中的数据
	// 	if num, ok := <-ch; ok {
	// 		fmt.Println("num=", num)
	// 	} else {
	// 		fmt.Println("channel关闭")
	// 		break
	// 	}
	// }
	//当管道关闭，此循环结束
	for num := range ch {
		fmt.Println("num=", num)
	}
}
