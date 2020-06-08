package testdemo

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeFunc(t *testing.T) {
	// TimeFunc01()
	fmt.Println(time.Unix(1588176000, 0).Format("2006-01-02 15:04:05"))
	t.Logf("test Add succ")
}

// 测试获取整点时间
func TestTimeFunc01(t *testing.T) {
	nowTime := time.Now().Local().Unix()
	// 获取凌晨时间
	bdStr := time.Unix(nowTime, 0).Format("2006-01-02")
	// 转时间
	bdTime, _ := time.ParseInLocation("2006-01-02", bdStr, time.Local)

	fmt.Println(nowTime)
	fmt.Println(bdTime.Unix())
	fmt.Println(time.Unix(nowTime, 0).Format("2006-01-02 15:04:05"))
	// 转时间
	bdTime1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-06-06 00:00:00", time.Local)
	fmt.Println(bdTime1.Unix())
}

// 1591372800
// 2020-06-06 15:45:26
