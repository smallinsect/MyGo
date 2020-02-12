package testdemo

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		status, err := conn.Write([]byte("爱白菜的小昆虫"))
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(status)
		time.Sleep(time.Second)
	}
}

func Userserver() {
	fmt.Println("server start ...")
	ser, err := net.Listen("tcp", ":7777")
	if err != nil {
		fmt.Println("server err ...")
		return
	}
	for {
		conn, err := ser.Accept()
		if err != nil {
			fmt.Println("new client connect err ...")
			continue
		}
		go handleConnection(conn)
	}
}
