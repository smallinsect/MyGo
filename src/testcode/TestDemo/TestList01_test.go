package testdemo

import (
	"container/list"
	"fmt"
	"runtime"
	"testing"
)

func TestListFunc(t *testing.T) {
	ListFunc()
	t.Logf("test ChannelFunc succ")
}

// 测试Go库函数的双向链表
func TestListFunc01(t *testing.T) {
	pli := list.New()

	elem1 := pli.PushBack("1111")
	elem2 := pli.PushFront("2222")
	elem3 := pli.PushBack("3333")
	fmt.Println(elem1)
	fmt.Println(elem2)
	fmt.Println(elem3)
	fmt.Println("===================")
	for e := pli.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
	pli.MoveToFront(elem3)
	fmt.Println("===================")
	for e := pli.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}
	fmt.Println("===================")
	fmt.Printf("%+v\n", elem3)
	pli.Remove(elem3)
	fmt.Println("===================")
	fmt.Printf("%+v\n", elem3)
	fmt.Println("===================")
	for e := pli.Front(); e != nil; e = e.Next() {
		fmt.Println(e)
	}

	t.Logf("test ChannelFunc succ")
}

// 测试Go库函数的双向链表的安全性
func TestListFunc02(t *testing.T) {
	runtime.GOMAXPROCS(8)
	l := list.New()
	ls := list.New()

	for i := 0; i < 10000; i++ {
		ls.PushBack(i)
	}
	go ls.Remove(l.Back())
	l.PushBackList(ls)
	t.Log("success ...")
}
