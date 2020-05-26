package testdemo

import (
	"fmt"
	"testing"
)

func TestUsemap(t *testing.T) {
	Usemap()
	t.Logf("test Add succ")
}

// 测试修改map中的指针值
func TestUsemap01(t *testing.T) {
	m := make(map[int64]*Person)
	m[111] = &Person{
		Uid:  1,
		Name: "1111",
		Age:  1,
	}

	fmt.Println(m[111])
	val, _ := m[111]
	val.Name = "2222"
	fmt.Println(m[111])

	t.Log("TestUsemap01 测试成功")
}
