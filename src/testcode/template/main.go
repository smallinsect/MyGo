package main

import (
	"fmt"
	"html/template"
	"os"
)

// Person struct
type Person struct {
	Name string
	Age  string
}

func main() {
	fmt.Println("http sever start ...")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{"Henry", "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an err:", err)
	}
}
