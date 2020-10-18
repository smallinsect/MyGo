package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type web struct{}

func (web) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	fmt.Println(request.URL.Path)
	if request.URL.Path == "/baidu" {
		client := http.DefaultClient
		resp, err := client.Get("http://www.baidu.com")
		if err != nil {
			writer.Write([]byte("请求失败"))
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		writer.Write(body)
		return
	}

	writer.Write([]byte("默认http服务器"))
}

func main() {

	c := make(chan os.Signal)

	http.ListenAndServe(":8888", web{})

	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println("服务器结束", s)
}
