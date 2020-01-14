// 15_设备文件的使用
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	//关闭后 不能往屏幕上写东西
	// os.Stdout.Close()

	//关闭后 不能从屏幕上读取数据
	// os.Stdin.Close()
	os.Stdout.WriteString("arr ")
}
