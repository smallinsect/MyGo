// 21_copy的使用
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	srcSlip := []int{1, 2}
	dstSlip := []int{6, 6, 6, 6, 6}

	fmt.Println(dstSlip)
	copy(dstSlip, srcSlip)
	fmt.Println(dstSlip)
}
