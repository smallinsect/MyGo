package main

import (
	"fmt"
	"log"
	"time"

	cron "gopkg.in/robfig/cron.v2"
)

func main1() {
	i := 0
	c := cron.New()
	// 也可以秒级任务
	// c := cron.New(cron.WithSeconds()) // 创建定时任务 秒
	//   spec := "0 */1 * * * *" // 每一分钟 如果用到分钟, 秒要设置为0
	//spec := "* */1 * * * *"
	fmt.Println(c)
	spec := "*/10 * * * * *" // 每一秒钟，
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
		fmt.Println("cron running:", i)
	})
	c.Start()
	select {}
}

func main2() {
	c := cron.New()

	now := time.Now()
	fmt.Println(now.UTC())
	fmt.Println(now)
	spec := "TZ=UTC */3 * 15 * * *" // 每一秒钟，
	_, err := c.AddFunc(spec, func() {
		now := time.Now()
		fmt.Println(now.UTC())
		fmt.Println(now)
	})
	if err != nil {
		log.Println(err)
	} else {
		c.Start()
		select {}
	}
}

var (
	layout    string = "2006-01-02 15:04:05"
	startTime string = "2020-03-25 10:00:00"
)

func main() {
	// Unix()时间戳，和时区无关。

	// 字符串转时间，默认UTC
	fmt.Println("==================================")
	start, _ := time.Parse(layout, startTime)
	fmt.Println(start)
	fmt.Println(start.Unix())
	fmt.Println(start.UTC())
	fmt.Println(start.UTC().Unix())

	// 字符串转时间，添加时区
	fmt.Println("==================================")
	startT, _ := time.ParseInLocation(layout, "2020-03-25 18:00:00", time.Local)
	fmt.Println(startT)
	fmt.Println(startT.Unix())
	fmt.Println(startT.UTC())
	fmt.Println(startT.UTC().Unix())

	// Now()获取当前系统时区当前时间
	fmt.Println("==================================")
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UTC())
	fmt.Println(now.UTC().Unix())

	// 现在时间，修改时区
	fmt.Println("==================================")
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fmt.Println(now.In(cstSh))
	var utcSh, _ = time.LoadLocation("UTC")
	fmt.Println(now.In(utcSh))

	fmt.Println("==================================")
	fmt.Println(time.Now().Zone())
	fmt.Println(time.LoadLocation("Local"))                                //获取时区
	timeStr := time.Now().Format("2006-01-02 15:04:05")                    //转化所需模板
	loc, _ := time.LoadLocation("Local")                                   //获取时区
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)) //使用模板在对应时区转化为time.time类型
	fmt.Println(time.Local)
}

// 凌晨时间
func ZeroTime(t time.Time) (zero_tm int64) {
	zero_tm = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tLocation).Unix()
	return
}
