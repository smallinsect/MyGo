package mypkg

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

// TestMapRoute is a function
func TestMapRoute() {
	m := make(map[int]int, 5)

	for i := 0; i < 2; i++ {
		go func(m map[int]int, val int) {
			lock.Lock()
			m[8] = val
			lock.Unlock()
		}(m, i)
	}

	time.Sleep(2 * time.Second)
	lock.Lock()
	fmt.Println(m)
	lock.Unlock()
}
