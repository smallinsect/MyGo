// 01_创建goroutine
package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is newTask")
		time.Sleep(time.Second)
	}
}

func TestGoroutine(t *testing.T) {
	fmt.Println("Hello World!")

	//创建协程并运行
	go newTask()

	for {
		fmt.Println("this is main")
		time.Sleep(time.Second)
	}
}
