package testchannel

import "fmt"

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
