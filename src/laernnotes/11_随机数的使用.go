// 11_随机数的使用
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	//设置随机种子
	rand.Seed(time.Now().UnixNano()) //以当前时间作为种子

	//产生随机数
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}
}
