package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

type ProxyHandler struct {
}

func (ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("1111111"))
}

func main() {
	c := make(chan os.Signal)

	http.ListenAndServe(":8080", ProxyHandler{})

	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println("服务器结束", s)
}
