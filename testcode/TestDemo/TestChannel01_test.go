package testdemo

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestChannelFunc(t *testing.T) {
	ChannelFunc()
	t.Log("test ChannelFunc succ")
}

func TestChannelFunc01(t *testing.T) {
	ch := make(chan int, 5)
	go func() {
		for i := range ch {
			fmt.Println("goroutine 1", i)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for i := range ch {
			fmt.Println("goroutine 2", i)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		i := 0
		for {
			ch <- i
			i++
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	t.Log("test ChannelFunc succ", sig)
}
