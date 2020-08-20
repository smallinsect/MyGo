pacakpackage main

import "fmt"

type Person struct {
	RankId int64
	Score  int64
	Name   string
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := make([]int, 0)

	arr2 = append(arr2, arr1[:10]...)
	fmt.Println(arr2)
}
