package testdemo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTestGoroutine(t *testing.T) {
	TestGoroutine()
	t.Log("test Add succ")
}

type value struct {
	mu    sync.Mutex
	value int
}

func TestTestGoroutine01(t *testing.T) {
	t.Log("test Add succ")

	var wg sync.WaitGroup

	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("Sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
	t.Log("test Add succ")
}
