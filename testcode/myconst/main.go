package main

import "fmt"

const (
	a = iota+1
	b
	c
	d
)

const (
	aa = iota
	bb
	cc
)

func main() {
	fmt.Println(a, b, c, d)
	fmt.Println(aa, bb, cc)
}
