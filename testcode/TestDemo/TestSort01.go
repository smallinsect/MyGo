package testdemo

import (
	"fmt"
	"sort"
)

type Group struct {
	Uid  int64 // 用户uid
	Hurt int64 //用户伤害
}

// SliceFunc ...
func SortFunc() {
	g := make([]*Group, 0)
	g = append(g, &Group{Uid: 111, Hurt: 1})
	g = append(g, &Group{Uid: 444, Hurt: 4})
	g = append(g, &Group{Uid: 555, Hurt: 4})
	g = append(g, &Group{Uid: 333, Hurt: 3})
	g = append(g, &Group{Uid: 222, Hurt: 2})

	for idx, val := range g {
		fmt.Println(idx, val)
	}

	fmt.Println("==========================")

	sort.Slice(g, func(i, j int) bool {
		return g[i].Hurt < g[j].Hurt
	})

	for idx, val := range g {
		fmt.Println(idx, val)
	}

}
