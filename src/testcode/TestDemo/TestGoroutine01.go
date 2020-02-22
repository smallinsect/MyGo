package testgoroutine

import (
	"fmt"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// TestGoroutine TestGoroutine
func TestGoroutine() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	fmt.Println("world")
	time.Sleep(time.Second)
}

func f1(i int) {
	defer wg.Done()
	fmt.Println(i)
}

var wg sync.WaitGroup

// TestWaitGroup TestWaitGroup
func TestWaitGroup() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}
