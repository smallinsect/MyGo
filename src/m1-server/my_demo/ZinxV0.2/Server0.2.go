package main

import "m1-server/zinx/znet"

// 服务端
func main() {
	//创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.2]")
	//启动server
	s.Serve()
}
