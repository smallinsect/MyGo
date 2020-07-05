package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

// Web1Handler Web1Handler
type web1Handler struct {
}

func (web1Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web1"))
}

// Web2Handler Web2Handler
type Web2Handler struct {
}

func (Web2Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web2"))
}

func main1() {

	c := make(chan os.Signal)

	go (func() {
		http.ListenAndServe("8888", web1Handler{})
	})()

	go (func() {
		http.ListenAndServe("88889", Web2Handler{})
	})()

	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println("服务器结束", s)
}
