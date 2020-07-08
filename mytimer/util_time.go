package main

import "time"

//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05" （go的诞生时间）
var TIME_LAYOUT string = "2006-01-02 15:04:05"

// 时区
var tLocation *time.Location

// 设置时区
func SetLocation(name string) (err error) {
	if name == "" {
		tLocation = time.Local // 获取当地时区
		return
	}
	tLocation, err = time.LoadLocation(name) // 设置时区，例："UTC"
	return
}

// 字符串转时间
func LayoutToTime(value string) (t time.Time, err error) {
	t, err = time.ParseInLocation(TIME_LAYOUT, value, tLocation)
	return
}

// 时间转字符串
func TimeToLayout(t time.Time) (layout string) {
	layout = t.Format(TIME_LAYOUT)
	return
}

// 凌晨时间，时区
func ZeroTime(t time.Time) (r int64) {
	r = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, tLocation).Unix()
	return
}

// 现在时间，时区
func TNow() (r time.Time) {
	r = time.Now().In(tLocation)
	return
}

// Unix时间戳转时间，时区
func TTime(t1, t2 int64) (r time.Time) {
	r = time.Unix(t1, t2).In(tLocation)
	return
}

// 同一天
func IsSameDay(t1, t2 int64) bool {
	y1, m1, d1 := TTime(t1, 0).Date()
	y2, m2, d2 := TTime(t2, 0).Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
