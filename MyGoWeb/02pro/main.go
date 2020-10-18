package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

type web1 struct{}

func (web1) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web1"))
}

type web2 struct{}

func (web2) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web2"))
}

func main() {

	c := make(chan os.Signal)

	go func() {
		http.ListenAndServe(":8888", web1{})
	}()

	go func() {
		http.ListenAndServe(":8889", web2{})
	}()

	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println("服务器结束", s)
}
