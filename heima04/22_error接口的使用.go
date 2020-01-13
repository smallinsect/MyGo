// 22_error接口的使用
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	err1 := fmt.Errorf("%s", "this is normal err1")
	fmt.Println(err1)

	err2 := errors.New("this is normal err2")
	fmt.Println(err2)
}
