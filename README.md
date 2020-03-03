# MyGo
## xine
# Go使用命令
go tool pprof 程序exe cpu_profile  
web 打开浏览器查看  
top----- 在profile 中输入top，会列出来几个最耗时的操作  
peek,list 妙用  
peek 是用来查询 函数名字的(这个名字是list需要使用的名字，并不完全等于函数名)  
list 是用来将函数时间消耗列出来的  
1）list main.main  
peek findMapMax (因为根据1可以看出来消耗都在 findMapMax)  
list main.findMapMax (根据2可以看出来名字是 main.findMapMax)  

# Go学习网址
安装包下载地址为：  
https://golang.org/dl/  
如果打不开可以使用这个地址：  
https://golang.google.cn/dl/  
Go 安装包下载  
https://studygolang.com/dl  
编辑器源码  
http://liteide.org/cn/  
编辑器  
https://sourceforge.net/projects/liteide/files/  
编辑器 
https://sourceforge.net/projects/liteide/  
go中文社区  
https://studygolang.com/  
go中文在线文档  
https://studygolang.com/pkgdoc  
json格式检查网站  
https://www.json.cn  
