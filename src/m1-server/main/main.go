// main.go
package main

import (
	"fmt"
	"m1-server/player"
	"math/rand"
	"time"
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
	// mypkg.TestMap()
	// mypkg.TestMapRoute()
	// var small mypkg.Person = &mypkg.Small{"111", "222"}

	// small.Say()
	// mypkg.TestTime()

	rand.Seed(time.Now().UnixNano())
	// for cnt := 0; cnt < 20; cnt++ {
	// 	remainLayerSum := int64(5)
	// 	var genStarNum int64 = 9             //需要生成的星星数量
	// 	arr := make([]int64, remainLayerSum) //每层星星数量
	// 	x := genStarNum / remainLayerSum     //平均星星数量
	// 	y := genStarNum % remainLayerSum     //多出的星星
	// 	//每层的星星数量大于1 每层扣一个星星 拿去随机
	// 	if x > 1 {
	// 		x = x - 1
	// 		y = y + remainLayerSum
	// 	}
	// 	//每层的基础星星数量
	// 	for idx := 0; idx < 5; idx++ {
	// 		arr[idx] = x
	// 	}
	// 	//将多出的星星拿来随机分配给每层
	// 	for i := int64(0); i < y; i++ {
	// 		idx := rand.Int() % 5
	// 		arr[idx]++
	// 	}
	// 	fmt.Println(arr)
	// }
	// var a int = rand.Intn(2)
	// fmt.Println(a)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Int31n(2))
	// 	fmt.Println(rand.Intn(2))
	// }
	// var min int = 22
	// var max int = 11
	// fmt.Println(min, max)
	// min, max = max, min
	// fmt.Println(min, max)

	// for cnt := 0; cnt < 100; cnt++ {
	// 	var stars int = 42
	// 	arrStar := make([]int, 5)
	// 	for i := 0; i < stars; i++ {
	// 		idx := rand.Int() % 5
	// 		arrStar[idx]++
	// 	}
	// 	fmt.Println(arrStar)
	// }
}
