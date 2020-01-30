package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("http client start ...")

	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Get data err:", err)
		return
	}
	fmt.Println(string(data))
}
