package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start client ...")
	conn, err := net.Dial("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println("Dial failed, err:", err)
		return
	}
	buf := make([]byte, 4096)
	for {
		cnt, err := conn.Read(buf)
		fmt.Println("cnt:", cnt)
		if err != nil {
			fmt.Println("Read failed, err:", err)
			return
		}
	}
}
