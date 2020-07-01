# MyGo

# 类型转换

字符串转bool

```
var b bool
b, _ = strconv.ParseBool("true")
b, _ = strconv.ParseBool("false")
```

bool转字符串

```
var b bool
str := strconv.FormatBool(b)
```

字符串转整形

```
i, err := strconv.Atoi("123456789")
```

整形转字符串

```
str := strconv.Itoa(123456789)
```

int64转字符串

```
// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
str := strconv.FormatInt(val, 10)
```

字符串转int64

```
func ParseInt(s string, base int, bitSize int) (i int64, err error)
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

```
// bool类型转字符串
```











