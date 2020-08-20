package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	testTime string = "2020-07-08 10:00:00"
)

func main() {
	// Unix()时间戳，和时区无关。

	// 字符串转时间，默认UTC
	fmt.Println("==================================")
	start, _ := time.Parse(TIME_LAYOUT, testTime)
	fmt.Println(start)
	fmt.Println(start.Unix())
	fmt.Println(start.UTC())
	fmt.Println(start.UTC().Unix())

	// 字符串转时间，添加时区
	fmt.Println("==================================")
	startT, _ := time.ParseInLocation(TIME_LAYOUT, "2020-08-19 00:00:00", time.Local)
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

	// 设置现在时间
	fmt.Println("==================================")
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fmt.Println(now.In(cstSh))
	var utcSh, _ = time.LoadLocation("UTC")
	fmt.Println(now.In(utcSh))

	// 时区
	fmt.Println("==================================")
	fmt.Println(time.Now().Zone())
	fmt.Println(time.LoadLocation("Local"))                                //获取时区
	timeStr := time.Now().Format("2006-01-02 15:04:05")                    //转化所需模板
	loc, _ := time.LoadLocation("Local")                                   //获取时区
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)) //使用模板在对应时区转化为time.time类型
	fmt.Println(time.Local)

	// Unix时间转时间
	fmt.Println("==================================")
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UTC())
	fmt.Println(time.Unix(now.Unix(), 0))           // Unix()转时间，默认当地时间
	fmt.Println(time.Unix(now.Unix(), 0).In(utcSh)) //Unix转时间，设置时区

	// 时间转字符串
	fmt.Println("==================================")
	lnow := now.Format(TIME_LAYOUT)
	fmt.Println(lnow)
	lnow = now.In(utcSh).Format(TIME_LAYOUT)
	fmt.Println(lnow)
}

func main3() {
	texp1 := "TZ=UTC */1 * 9 * * *"
	t.AddFunc(texp1, func() {
		fmt.Println("1", time.Now().UTC())
		fmt.Println("2", time.Now())
	})
	texp2 := "*/5 * * * * *"
	t.AddFunc(texp2, func() {
		fmt.Println("1", time.Now().UTC())
		fmt.Println("2", time.Now())
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	fmt.Println("server closing down signal:", sig)
}
