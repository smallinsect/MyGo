package testchannel

import (
	"fmt"
	"sync"
	"time"
)

//ChannelFunc ChannelFunc
func ChannelFunc() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		fmt.Println("c end ....")
	}()
	go func() {
		for i := 0; i < 10; i++ {
			x := <-c
			fmt.Println(x)
		}
		fmt.Println("x end ....")
	}()
}

func ChannelFunc01() {
	ch := make(chan int, 50)
	wg := new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		ch <- i
		go func(val int) {
			defer func() {
				wg.Done()
				x := <-ch
				fmt.Println("end x=", x)
			}()
			fmt.Println("start i=", val)
			time.Sleep(time.Second)
		}(i)
	}
}
