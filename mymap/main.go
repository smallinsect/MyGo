package main

import "fmt"

type Person struct {
	RankId int64
	Score  int64
	Name   string
}

func main() {
	m := make(map[string]int64)
	m["aaa"] = 111
	m["bbb"] = 222
	m["ccc"] = 333
	m["ddd"] = 444
	for key := range m {
		fmt.Println(key)
	}
}
