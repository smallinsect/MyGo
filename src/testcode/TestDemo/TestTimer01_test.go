package testdemo

import (
	"fmt"
	"testing"
	"time"
)

func f() {
	fmt.Println("f() run ...")
}

func TestTimer(t *testing.T) {
	// 每1秒执行一次
	c := time.Tick(1 * time.Second)
	for {
		<-c
		go f()
	}
	t.Logf("test Add succ")
}

func TestTimer01(t *testing.T) {
	var ch chan int
	// 定时任务
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		ch <- 1
	}()
	<-ch
	t.Logf("test Add succ")
}
