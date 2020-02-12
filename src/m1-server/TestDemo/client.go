package testdemo

import (
	"fmt"
	"net"
	"time"
)

func Useclient() {
	fmt.Println("client start ...")
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client err ...")
		return
	}
	for {
		var buf [512]byte
		status, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println("bufio.NewReader(conn).ReadString err ...")
			break
		}
		fmt.Println(string(buf[0:status]))
		time.Sleep(time.Second)
	}
	conn.Close()
}
