package testdemo

import (
	"fmt"
	"testing"
	"time"
)

var TIME_LAYOUT string = "2006-01-02 15:04:05"

func f() {
	fmt.Println("f() run ...")
}

func TestTimer(t *testing.T) {
	// 每1秒执行一次
	c := time.Tick(1 * time.Second)
	for {
		<-c
		go f()
	}
	t.Logf("test Add succ")
}

func TestTimer01(t *testing.T) {
	var ch chan int
	// 定时任务
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		ch <- 1
	}()
	<-ch
	t.Logf("test Add succ")
}

func TestTimer02(t *testing.T) {
	// 设置时区，如果name是""或"UTC"，返回UTC；
	// 如果name是"Local"，返回Local；
	// 否则name应该是IANA时区数据库里有记录的地点名（该数据库记录了地点和对应的时区），如"America/New_York"。
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}

	// 创建时间，依次是：年、月、日、时、分、秒、纳秒(1秒＝1000毫秒＝1000000微秒＝1000000000纳秒)、时区
	t1 := time.Date(2018, 7, 7, 12, 12, 12, 500000000, location)
	fmt.Println(t1) // 2018-07-07 12:12:12.5 -0400 EDT

	// 将字符串转成时间，时间格式字符串："2006-01-02 15:04:05"（Go语言规定）
	t2, err := time.Parse("2006-01-02 15:04:05", "2018-07-07 09:10:05")
	fmt.Println(t2) // 2018-07-07 09:10:05 +0000 UTC

	// 将字符串转成时间，需要传入时区
	t3, err := time.ParseInLocation("20060102", "20180707", time.UTC)
	fmt.Println(t3) // 2018-07-07 00:00:00 +0000 UTC

	// 获取当前时间
	t4 := time.Now()
	fmt.Println(t4) // 2018-05-23 20:50:08.9873106 +0800 CST m=+0.012895501

	// 格式化输出
	fmt.Println(t4.Format("2006-01-02 15:04:05")) // 2018-05-23 20:50:08
	fmt.Println(t4.Format("02/01/2006 15:04:05")) // 23/05/2018 20:50:08
	fmt.Println(t4.Format("2006-01-02"))          // 2018-05-23
	fmt.Println(t4.Format("15:04:05"))            // 20:50:08
	fmt.Println(t4.Format("January 2,2006"))      // May 23,2018

	// 获取世界统一时间
	t5 := t4.UTC()
	fmt.Println(t5) // 2018-05-23 12:50:08.9873106 +0000 UTC

	// 获取本地时间
	t6 := t5.Local()
	fmt.Println(t6) // 2018-05-23 20:50:08.9873106 +0800 CST

	// 获取指定时区的时间
	t7 := t6.In(location)
	fmt.Println(t7) // 2018-05-23 08:50:08.9873106 -0400 EDT

	// 获取Unix时间戳，单位：秒，即从时间点1970-01-01 00:00:00 UTC到时间点t所经过的时间
	timestamp := t7.Unix()
	fmt.Println(timestamp) // 1527080185

	// 获取Unix时间戳，单位：纳秒，常用于作为rand的随机数种子
	timestamp = t7.UnixNano()
	fmt.Println(timestamp) // 1527080185738346000

	// 判断两个时间是否相等，会判断时区等信息，不同时区也可以用此进行比较
	fmt.Println(t7.Equal(t6)) // true

	// 判断t4是否在t3之前
	fmt.Println(t4.Before(t3)) // true

	// 判断t4是否在t3之后
	fmt.Println(t4.After(t3)) // false

	// 返回时间的年、月、日
	y, m, d := t4.Date()
	fmt.Printf("年：%d，月：%d，日：%d\n", y, m, d) // 年：2018，月：5，日：23

	// 返回时间的时、分、秒
	h, minute, s := t4.Clock()
	fmt.Printf("时：%d，分：%d，秒：%d\n", h, minute, s) // 时：21，分：5，秒：41

	// 单独获取年、月、日、时、分、秒、星期
	fmt.Printf("年：%d，月：%d，日：%d，时：%d，分：%d，秒：%d，星期：%d\n", t4.Year(), t4.Month(), t4.Day(), t4.Hour(), t4.Minute(), t4.Second(), t4.Weekday()) // 年：2018，月：5，日：23，时：21，分：9，秒：56，星期：3

	t8, err := time.Parse("2006-01-02 15:04:05", "2018-01-01 00:00:00")
	// 增加100秒，time.Duration是以纳秒为单位，time.Second=1000 000 000。参数可以为负数就是减少
	t9 := t8.Add(time.Duration(100) * time.Second)
	fmt.Println(t9) // 2018-01-01 00:01:40 +0000 UTC

	// 增加或减少年、月、日
	t10 := t8.AddDate(1, 1, -1)
	fmt.Println(t10) // 2019-01-31 00:00:00 +0000 UTC

	// 计算两个时间之间的差
	dur := t8.Sub(t9)
	fmt.Println(dur.Seconds()) // -100

	// 将时间戳转成Time对象，第一个参数为秒，第二个参数为纳秒，如果传秒数，则纳秒传入0；如果传入纳秒，则秒传入0
	t11 := time.Unix(0, timestamp)
	fmt.Println(t11.Format("2006-01-02 15:04:05"))

	//=======================================================================定时器、休眠等
	// Timer，单次时间事件，指定时间后向通道C发送当时时间
	timer := time.NewTimer(time.Duration(1) * time.Second)
	fmt.Println(<-timer.C)

	// 也可配合select使用
	timer = time.NewTimer(time.Duration(1) * time.Second)
	select {
	case <-timer.C:
		fmt.Println("执行...")
	}

	// 用Timer实现定时器
	timer = time.NewTimer(time.Duration(1) * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("Timer定时器...")
			timer.Reset(time.Duration(1) * time.Second) // 重新开始计时
		}
	}

	// 开启一个新协程，在指定时间后执行给定函数，所以测试时，需要将主协程休眠几秒才能看到执行结果
	time.AfterFunc(time.Duration(1)*time.Second, func() {
		fmt.Println("AfterFunc...")
	})
	// 当前协程休眠指定时间
	time.Sleep(2 * time.Second)

	// 指定时间后向通道C发送当时时间
	tt := <-time.After(time.Duration(1) * time.Second)
	fmt.Println(tt)

	// Ticker保管一个通道，并每隔一段时间向其传递"tick"。
	ticker := time.NewTicker(time.Duration(1) * time.Second)
	// 用Ticker实现定时器
	for {
		select {
		case <-ticker.C:
			fmt.Println("Ticker...")
		}
	}
}

func TestTimer03(t *testing.T) {
	// 将字符串转成时间，时间格式字符串："2006-01-02 15:04:05"（Go语言规定）
	stime, _ := time.Parse("2006-01-02 15:04:05", "2020-04-08 10:38:00")
	ntime := time.Now()
	// 指定时间执行
	st := stime.UnixNano() - ntime.UnixNano()
	fmt.Println(st)
	fmt.Println(stime.Format("2006-01-02 15:04:05"))
	fmt.Println(ntime.Format("2006-01-02 15:04:05"))
	fmt.Println(time.Duration(st))
	timer := time.NewTimer(time.Duration(st))
	for {
		fmt.Println(111)
		<-timer.C
		// 时间转字符串格式
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println(222)
	}
	t.Log("suc ...")
}

func TestTimer04(t *testing.T) {
	cstZone := time.FixedZone("CST", 8*3600) // 东八
	// var cstZone, _ = time.LoadLocation("Asia/Shanghai") //上海
	ntime := time.Now()
	stime, _ := time.ParseInLocation(TIME_LAYOUT, "2020-04-08 11:59:00", cstZone)
	fmt.Println(stime.Sub(ntime.In(cstZone)))
	timer := time.AfterFunc(stime.Sub(ntime), func() {
		fmt.Println("time.AfterFunc ...")
	})

	// <-timer.C
	timer.Stop()
	timer.Reset(1 * time.Second)
	fmt.Println("timer.Stop()")
	for {

	}
	t.Log("suc ...")
}

func TestTimer05(t *testing.T) {
	ti := time.Now()
	cstZone := time.FixedZone("CST", 8*3600) // 东八
	// fmt.Println(ti.In(cstZone).Format("01-02-2006 15:04:05"))
	// 字符串转换时间，添加时区。
	stime, _ := time.ParseInLocation(TIME_LAYOUT, "2020-04-08 11:36:00", cstZone)
	fmt.Println(ti.Format(TIME_LAYOUT), ti.Unix())
	fmt.Println(stime.Format(TIME_LAYOUT), stime.Unix())
	fmt.Println(ti.Unix() - stime.Unix())
	t.Log("suc ...")
}

// 测试定时器执行的函数，重新创建，是否失效
func TestTimer06(t *testing.T) {

	fmt.Println("0", time.Now().Format(TIME_LAYOUT))
	ch := make(chan int)
	timer := time.AfterFunc(3*time.Second, func() {
		fmt.Println("timer 1", time.Now().Format(TIME_LAYOUT))
		ch <- 1
	})
	time.Sleep(1 * time.Second)
	timer = time.AfterFunc(4*time.Second, func() {
		fmt.Println("timer 2", time.Now().Format(TIME_LAYOUT))
		ch <- 2
	})
	timer.Reset(5 * time.Second)
	fmt.Println("3", time.Now().Format(TIME_LAYOUT))

	select {
	case <-ch:
		fmt.Println("4", time.Now().Format(TIME_LAYOUT))
		break
	}
	t.Log("suc ...")
}
