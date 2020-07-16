package main

import (
	"fmt"
)

func main() {
	t := 4
	switch t {
	case 1:
		fmt.Println("小昆虫1")
	case 2:
		fmt.Println("小昆虫2")
	case 3, 4:
		fmt.Println("小昆虫3")
	}
}
