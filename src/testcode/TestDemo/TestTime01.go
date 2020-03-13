package testdemo

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

//TimeFunc01 TimeFunc01
func TimeFunc01() {
	t := int64(1546926630)      //外部传入的时间戳（秒为单位），必须为int64类型
	t1 := "2020-02-20 18:41:00" //外部传入的时间字符串

	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	timeTemplate2 := "2006/01/02 15:04:05" //其他类型
	timeTemplate3 := "2006-01-02"          //其他类型
	timeTemplate4 := "15:04:05"            //其他类型

	// ======= 将时间戳格式化为日期字符串 =======
	fmt.Println(time.Unix(t, 0).Format(timeTemplate1)) //输出：2019-01-08 13:50:30
	fmt.Println(time.Unix(t, 0).Format(timeTemplate2)) //输出：2019/01/08 13:50:30
	fmt.Println(time.Unix(t, 0).Format(timeTemplate3)) //输出：2019-01-08
	fmt.Println(time.Unix(t, 0).Format(timeTemplate4)) //输出：13:50:30

	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	fmt.Println(stamp.Unix())                                       //输出：1546926630
}

//TimeFunc02 当前时间加1个小时
func TimeFunc02() {
	nTime := time.Now() //现在时间
	fmt.Println(nTime)
	fmt.Println(nTime.Unix())                                 //当前时间戳
	fmt.Println(nTime.Add(time.Minute).Unix())                //当前时间戳+一分钟
	fmt.Println(nTime.Add(time.Minute).Unix() - nTime.Unix()) //时差一分钟
	fmt.Println(nTime.Add(time.Hour).Unix())                  //当前时间戳+一小时
	fmt.Println(nTime.Add(time.Hour).Unix() - nTime.Unix())   //时差一小时
}
