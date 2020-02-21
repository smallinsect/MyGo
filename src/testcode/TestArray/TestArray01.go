package testarray

import "fmt"

//ArrayFunc ArrayFunc
func ArrayFunc() {
	arr := make([]int, 0)
	fmt.Println(len(arr))
	fmt.Println(arr)
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Println(len(arr))
	fmt.Println(arr)
}

// User User
type User struct {
	Name string
}

// ArrayFunc01 ArrayFunc01
func ArrayFunc01() {
	mus := make([]*User, 0)
	var matchNum int = 1
	if len(mus) == 0 {
		mus = getMatchList()
		matchNum++
	}
	fmt.Println(mus)
	fmt.Println(matchNum)
}

func getMatchList() (r []*User) {
	r = make([]*User, 0)
	return
}
