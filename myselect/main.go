package main

import (
	"fmt"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpConn := conn.(*net.TCPConn)
	tcpConn.SetReadBuffer(1024 * 64)
	tcpConn.SetReadBuffer(1024 * 64)
}
