package mypkg

import (
	"fmt"
	"time"
)

// TestGroute1 is a function
func TestGroute1() {
	for i := 0; i < 5; i++ {
		fmt.Println("i=", i)
	}
	time.Sleep(time.Second)
}

// TestPipe is a function
func TestPipe() {
	pipe := make(chan int32, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	// pipe <- 4
	var t1 int32
	t1 = <-pipe

	fmt.Println(len(pipe), pipe)
	fmt.Println(t1)
}

// TestMap is a function
func TestMap() {
	m := make(map[int]int, 5)

	m[5] = 55
	m[1] = 11
	m[3] = 33
	m[2] = 22
	m[4] = 44

	for _, val := range m {
		fmt.Println(val)
	}
}
