package testdemo

import (
	"fmt"
	"sync"
)

// MapFunc ...
func MapFunc() {
	var m sync.Map
	m.Store("1111", &Person{Name: "aaaaa"})
	m.Store("2222", &Person{Name: "bbbbb"})
	m.Store("3333", &Person{Name: "ccccc"})

	p, _ := m.Load("1111")
	// fmt.Println(m)
	fmt.Println(p)
	m.Delete("1111")
	// fmt.Println(m)
	fmt.Println(p)

	// m := make(map[string]*Person)
	// m["1111"] = &Person{Name: "aaaaa"}
	// m["2222"] = &Person{Name: "bbbbb"}
	// m["3333"] = &Person{Name: "ccccc"}

	// p := m["1111"]
	// fmt.Println(p)
	// fmt.Println(m)
	// delete(m, "1111")
	// fmt.Println(p)
	// fmt.Println(m)
}

// MapFunc01 ...
func MapFunc01() {

}
