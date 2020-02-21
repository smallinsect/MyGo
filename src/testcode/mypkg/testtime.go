package mypkg

import (
	"fmt"
	"time"
)

// TestTime is
func TestTime() {
	fmt.Println("bbb")
	for {
		select {
		case <-time.After(time.Millisecond):
			fmt.Println("aaa")
		}
	}
}
