package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"runtime/pprof"
	"syscall"
	"time"
)

func init() {
	fmt.Println("我是init函数")
	fmt.Println("在main函数执行之前的init函数")
}

type Extend struct {
	Tid int64 `json:"tid"`
}

func main() {
	exten := "{\"Tid\": 10}"
	var extenT Extend
	err := json.Unmarshal([]byte(exten), &extenT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(extenT)
	fmt.Println(exten)
}

type Cat struct {
}

func (this *Cat) LoginRequest(req *JumpRequest) (res *JumpResponse) {
	fmt.Printf("%+v\n", req)
	fmt.Println("喵喵....")
	res = &JumpResponse{
		NameRes: "小昆虫",
		AgeRes:  111,
	}
	return
}

type JumpRequest struct {
	NameReq string
	AgeReq  int64
}

type JumpResponse struct {
	NameRes string
	AgeRes  int64
}

// GenerateGroupNum GenerateGroupNum
func GenerateGroupNum(arrLen int, genNum int64) (arrNum []int64) {
	if genNum <= 0 {
		return
	}
	arrNum = make([]int64, arrLen)
	avgNum := genNum / int64(arrLen) //平均的数量
	spsNum := genNum % int64(arrLen) //剩余的数量
	//平均的数量大于1 扣除1个 拿去随机
	if avgNum > 1 {
		avgNum = avgNum - 1
		spsNum = spsNum + int64(arrLen)
	}
	//数组的基础数量
	for idx := 0; idx < arrLen; idx++ {
		arrNum[idx] = avgNum
	}
	//将未分配的数 随机到数组中
	for i := int64(0); i < spsNum; i++ {
		// idx := utils.RandomInt(0, arrLen-1)
		idx := rand.Intn(arrLen)
		arrNum[idx]++
	}
	return
}

// Demo ...
func Demo() {
	//添加记录CPU使用率的文件--------------------
	cpuf, err1 := os.Create("cpu_profile")
	if err1 != nil {
		log.Fatal(err1)
	}
	pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()
	//添加记录CPU使用率的文件--------------------
	fmt.Println("Hello World!")
	// var p player.Player

	// fmt.Println(p)
	// p.SetInfo("爱白菜1", "小昆虫1")
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

	// for i := 0; i < 100; i++ {
	// 	arr := GenerateGroupNum(5, 42)
	// 	fmt.Println(arr)
	// }
	// fmt.Printf("%d", time.Minute)
	// for cnt := 0; cnt < 100; cnt++ {
	// 	arr := make([]int, 7)
	// 	for i := 0; i < 5; i++ {
	// 		idx := rand.Int() % 7
	// 		arr[idx]++
	// 	}
	// 	fmt.Println(arr)
	// }
	// testprofile.ProfileFunc()
	// testprofile.ProfileFunc01()
	// testgoroutine.TestGoroutine()
	// testgoroutine.TestWaitGroup()
	// testchannel.ChannelFunc01()

	// testlist.ListFunc()
	// testtime.TimeFunc01()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Println("server closing down signal:", sig)
	fmt.Printf("server closing down (signal: %v)\n", sig)
}

func IFunc() {
	fmt.Println("主函数")
	req := &JumpRequest{
		NameReq: "小白菜",
		AgeReq:  222,
	}
	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf(req)

	cat := new(Cat)
	t := reflect.ValueOf(cat)
	method := t.MethodByName("LoginRequest")
	if method.IsValid() {
		fmt.Println("有效方法...")
		res := method.Call(in)
		for _, item := range res {
			fmt.Printf("%+v\n", item)
		}
		fmt.Println("结束...")
	} else {
		fmt.Println("无效方法...")
	}
}
