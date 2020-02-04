package main

import (
	"fmt"
	"net"
	"time"
)

//客户端
func main() {
	fmt.Println("client start ...")
	time.Sleep(1 * time.Second)
	//1 直接连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client start err exit", err)
		return
	}
	//2 链接调用Write写数据
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.1 ..."))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error", err)
			return
		}
		fmt.Println("server call back", buf[:cnt], cnt)
		fmt.Println("server call back", string(buf[:cnt]), cnt)
		time.Sleep(1 * time.Second)
	}
}
