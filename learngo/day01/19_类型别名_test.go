// 19_类型别名
package main

import (
	"fmt"
	"testing"
)

func TestNoVar(t *testing.T) {
	//给int64起一个别名
	type bigint int64

	var a bigint

	fmt.Printf("a type %T\n", a)

	// type(
	// 	i int32
	// 	char type
	// )
}
