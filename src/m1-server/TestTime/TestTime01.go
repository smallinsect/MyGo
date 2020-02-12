package testtime

import (
	"fmt"
	"time"
)

//TimeFunc TimeFunc
func TimeFunc() {
	t := time.Now()
	fmt.Println("t", t)
	fmt.Println("t.Unix()", t.Unix())
	fmt.Println("t.Add(1)", t.Add(1))
	fmt.Println("t.Add(1).Unix()", t.Add(1).Unix())
	fmt.Println(time.Minute)
}
