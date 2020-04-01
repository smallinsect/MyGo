package testdemo

import (
	"fmt"
	"testing"
)

// 测试排序
func TestSortFunc(t *testing.T) {
	SortFunc()
	t.Logf("test Add succ")
}

// 测试除数和求模
func TestSortFunc01(t *testing.T) {
	fmt.Println(120 / 100)
	fmt.Println(120 % 100)
	fmt.Println(100 % 100)
	t.Log("TestSortFunc01")
}
