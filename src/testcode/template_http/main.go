package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var MyTemplate *template.Template

// Person struct
type Person struct {
	Name string
	Age  string
}

func main() {
	fmt.Println("server start ...")
	InitTemplate("index.html")
	http.HandleFunc("/user/info", UserInfo)
	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println("http listen failed err:", err)
	}
}

// InitTemplate InitTemplate
func InitTemplate(filename string) (err error) {
	MyTemplate, err = template.ParseFiles(filename)
	if err != nil {
		fmt.Println("template.ParseFiles err:", err)
		return
	}
	return
}

// UserInfo UserInfo
func UserInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "......")
	p := Person{"Henry", "31"}
	MyTemplate.Execute(w, p)
}
