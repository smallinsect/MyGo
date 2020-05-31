# MyGo
# Go学习网址
1. 安装包下载地址为：

* https://golang.org/dl/

2. 如果打不开可以使用这个地址：
```
https://golang.google.cn/dl/
```
3. Go 安装包下载
```
https://studygolang.com/dl
```
4. 编辑器源码
    * http://liteide.org/cn/  
5. 编辑器
    * https://sourceforge.net/projects/liteide/files/
6. 编辑器
    * https://sourceforge.net/projects/liteide/
7. go中文社区
    * https://studygolang.com/
8. go中文在线文档
    * https://studygolang.com/pkgdoc
9. json格式检查网站
    * https://www.json.cn
10. 教程
```
https://www.w3cschool.cn/go_internals/go_internals-dea92830.html
```
下载包的地址
```
https://github.com/golang/go
https://github.com/golang/text

https://github.com/golang/time.git


```



# Linux安装Go

下载链接

```
安装包下载地址为：https://golang.org/dl/
如果打不开可以使用这个地址：https://golang.google.cn/dl/
```

下载

```shell
# 在 ~ 下创建 go 文件夹，并进入 go 文件夹
mkdir ~/go && cd ~/go
下载的 go 压缩包
wget https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
```

解压

```
tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz
```

添加环境变量

```
# 习惯用vim，没有的话可以用命令`sudo apt-get install vim`安装一个
vim /etc/profile
# 在最后一行添加
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
# 保存退出后source一下（vim 的使用方法可以自己搜索一下）
source /etc/profile
```



# Gin框架搭建

官网文档

```
https://gin-gonic.com/docs/quickstart/
```

开启Go Module

```
go env -w GO111MODULE=on
```

开启代理

```
go env -w GOPROXY=https://goproxy.cn,direct
```

查看环境变量修改

```
go env
```

初始化项目模块

```
go mod init oceanlearn.teach/ginessential
```

下载Gin的依赖

```
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
```

添加MySQL驱动

```
go get github.com/go-sql-driver/mysql
```











# Go框架

Golang服务端开发之道
https://www.cnblogs.com/isaiah/
https://www.cnblogs.com/isaiah/p/7259036.html





# Go常用命令
1. 测试单个文件，一定要带上被测试的原文件
    * go test -v slice_test.go slice.go
2. 测试单个方法
    * go test -v -test.run TestFunc01

# Go使用性能命令
1. go tool pprof 程序exe cpu_profile
    * web 打开浏览器查看
    * top----- 在profile 中输入top，会列出来几个最耗时的操作
    * peek,list 妙用
    * peek 是用来查询 函数名字的(这个名字是list需要使用的名字，并不完全等于函数名)  
    * list 是用来将函数时间消耗列出来的
    * list main.main  
    * peek findMapMax (因为根据1可以看出来消耗都在 findMapMax)  
    * list main.findMapMax (根据2可以看出来名字是 main.findMapMax)  

# Go常用代码注意问题
1. 定时器
    * t := time.AfterFunc的t.C是nil值
2. cron 定时任务
    * 匹配符号
        - 星号(*):`表示 cron 表达式能匹配该字段的所有值。如在第5个字段使用星号(month)，表示每个月`
        - 斜线(/):`表示增长间隔，如第2个字段(minutes) 值是 3-59/15，表示每小时的第3分钟开始执行一次，之后 每隔 15 分钟执行一次（即 3（3+0*15）、18（3+1*15）、33（3+2*15）、48（3+3*15） 这些时间点执行），这里也可以表示为：3/15`
        - 逗号(,):`用于枚举值，如第6个字段值是 MON,WED,FRI，表示 星期一、三、五 执行`
        - 连字号(-):`表示一个范围，如第3个字段的值为 9-17 表示 9am 到 5pm 直接每个小时（包括9和17）`
        - 问号(?): `只用于日(Day of month) 和 星期(Day of week)，表示不指定值，可以用于代替*`
    * 样例
        - 比如我们的手机卡假设都是在每个月的开始时间就更新资费：
        - "0 0 0 1 * *" // 表示每个月1号的00:00:00
        - "0 1 1 1 * *" // 表示每个月1号的01:01:00
        - 每隔5秒执行一次："*/5 * * * * ?"
        - 每隔1分钟执行一次："0 */1 * * * ?"
        - 每天23点执行一次："0 0 23 * * ?"
        - 每天凌晨1点执行一次："0 0 1 * * ?"
        - 每月1号凌晨1点执行一次："0 0 1 1 * ?"
        - 在26分、29分、33分执行一次："0 26,29,33 * * * ?"
        - 每天的0点、13点、18点、21点都执行一次："0 0 0,13,18,21 * * ?"



# Go代码

```
//	若使用Marshal则不会格式化，这样不方便阅读
bytes, err := json.MarshalIndent(obj, "", " ")
```





































