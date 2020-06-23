# MyGo

# 类型转换

整形转字符串

```
// val是整形，10是转十进制
str := strconv.FormatInt(val, 10)
```

浮点数转字符串

```
// val是浮点数
str := strconv.FormatFloat(val, 'f', -1, 64)
```

字符串转整形

```
// 转10进制，64位整形
id, _ := strconv.ParseInt(strId, 10, 64)
```











