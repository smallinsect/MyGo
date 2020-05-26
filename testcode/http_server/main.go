package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http sever start ...")
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", Login)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

// Hello Hello
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "hello")
}

// Login Login
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "login")
}
