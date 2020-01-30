package mypkg

import (
	"fmt"
)

// Person is type
type Person interface {
	Say()
}

// Small is type
type Small struct {
	Name string
	Pass string
}

// Say is func
func (small *Small) Say() {
	fmt.Println(small.Name, small.Pass)
}
