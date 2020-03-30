// 16_文件的书写
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	//使用完毕，需要关闭文件
	defer f.Close()

	for i := 0; i < 10; i++ {
		buf := fmt.Sprintf("i = %d\n", i)
		f.WriteString(buf)
	}
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //2k大小

	//n代表读取的长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:n]))
}

//每次读取一行
func ReadLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区
	r := bufio.NewReader(f)

	for {
		//以'\n'读取数据
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件末尾
				break
			}
			fmt.Println(err)
		}
		fmt.Println(string(buf))
	}
}

func main() {
	fmt.Println("Hello World!")
	path := "./demo.txt"

	// WriteFile(path)
	// ReadFile(path)
	ReadLine(path)
}
