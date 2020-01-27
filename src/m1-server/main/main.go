// main.go
package main

import (
	"fmt"
	"m1-server/mypkg"
	"m1-server/player"
)

func main() {
	fmt.Println("Hello World!")
	var p player.Player

	// fmt.Println(p)
	p.SetInfo("爱白菜1", "小昆虫1")
	// fmt.Println(p)

	// mypkg.test_goroute1()
	// go mypkg.TestGroute1()
	// go mypkg.TestGroute1()
	// go mypkg.TestGroute1()
	// mypkg.TestPipe()
	mypkg.TestMap()
	// time.Sleep(5 * time.Second)
}
