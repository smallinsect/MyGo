package znet

import (
	"fmt"
	"m1-server/zinx/ziface"
	"net"
)

// Server 的接口实现，定义一个Server的服务器模块
type Server struct {
	Name      string //服务器的名称
	IPVersion string //服务器绑定的ip版本
	IP        string //服务器监听的IP
	Port      int    //服务器监听的端口
}

// Start Start
func (s *Server) Start() {
	fmt.Printf("[Start Server Listenner at IP:%s, Port:%d ...]\n", s.IP, s.Port)
	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error:", err)
			return
		}
		//2 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err:", err)
			return
		}
		fmt.Println("start Zinx server succ, ", s.Name, " succ, Listenning ...")
		//3 阻塞的等待客户端连接，处理客户端链接业务（读写）
		for {
			//如果有客户端连接过来
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			//客户端已经链接，做一些业务，做一个最基本的512字节长度
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Println("client:", buf[:cnt])
					//回写功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()
}

// Stop Stop
func (s *Server) Stop() {
	//讲一些服务器的资源、状态或者一些
}

// Serve Serve
func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()
	//阻塞状态
	select {}
}

/*
NewServer 初始化
*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8888,
	}
	return s
}
