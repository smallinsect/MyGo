package testlist

import (
	// "container/list"
	"fmt"
	"sync"
)

type Person struct {
	Uid  int32
	Name string
}

//ChannelFunc ChannelFunc
func ListFunc() {
	// var l list.List
	var m sync.Map
	m.Store(111, &Person{Uid: 111, Name: "小昆虫1"})
	m.Store(222, &Person{Uid: 222, Name: "小昆虫2"})
	m.Store(333, &Person{Uid: 333, Name: "小昆虫3"})
	m.Range(func(key, val interface{}) bool {
		fmt.Println(key, val)
		return true
	})
	if v, ok := m.Load(222); ok {
		v1 := v.(*Person)
		v1.Name = "小昆虫4"
	}
	m.Range(func(key, val interface{}) bool {
		fmt.Println(key, val)
		return true
	})

	// for i := 0; i < 10000; i++ {
	// 	e := new(list.Element)
	// 	e.Value = i
	// 	l.PushBack(e)
	// 	m.Store(i, e)
	// }

	// for i := 10000; i > 0; i-- {
	// 	if val, ok := m.Load(i); ok {
	// 		elem := val.(*list.Element)
	// 		l.MoveToFront(elem)
	// 	}
	// }
	// for item := l.Front(); item != nil; item = item.Next() {
	// 	fmt.Println(item.Value)
	// }
}
