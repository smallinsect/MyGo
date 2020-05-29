package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("heheh")
	reps, err := http.Get("")
	if err != nil {
		fmt.Println(err)
		return
	}
}
