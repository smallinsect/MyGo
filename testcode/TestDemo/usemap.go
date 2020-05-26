package testdemo

import "fmt"

// Usemap ...
func Usemap() {
	m := make(map[string]Person)
	m["1111"] = Person{Name: "小昆虫1", Age: 1}
	m["2222"] = Person{Name: "小昆虫2", Age: 2}
	m["3333"] = Person{Name: "小昆虫3", Age: 3}
	m["1111"] = Person{Name: "小昆虫4", Age: 4}
	fmt.Printf("%v\n", m)
	delete(m, "2222")
	fmt.Printf("%v\n", m)
}
