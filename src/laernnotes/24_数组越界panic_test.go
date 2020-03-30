// 24_数组越界panic
package main

import (
	"fmt"
)

func at(idx int) {
	var arr [10]int
	arr[idx] = 100
}

func main() {
	fmt.Println("Hello World!")
	at(111)
}
