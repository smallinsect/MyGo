package main

import (
	"fmt"
)

var Upper []string = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func main() {
	// PlayerCount := int64(0)

	// wc := new(sync.WaitGroup)
	// wc.Add(10000)
	// for i := 0; i < 10000; i++ {
	// 	go func() {
	// 		atomic.AddInt64(&PlayerCount, 1)
	// 		fmt.Println(PlayerCount)
	// 		wc.Done()
	// 	}()
	// }

	// wc.Wait()
	count := int64(269989999)
	countT := count + 10000
	left := countT / 100000000
	right := countT % 100000000

	r := ""
	if left > 0 {
		left = countT / 10000000
		upper := Upper[left-10]
		r = fmt.Sprintf("%s%07d", upper, right)
	} else {
		r = fmt.Sprintf("%d", right)
	}
	fmt.Println(r)
}
