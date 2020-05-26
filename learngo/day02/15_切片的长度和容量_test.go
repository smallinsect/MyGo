// 15_切片的长度和容量
package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	fmt.Println("Hello World!")

	a := []int{1, 2, 3, 4, 5, 6, 7}
	s := a[0:3:5]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = a[1:4:5]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

}
