package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("heheh")
	reps, err := http.Get("https://www.runoob.com/linux/linux-vim.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reps.Body.Close()

	if reps.StatusCode != http.StatusOK {
		fmt.Println("", reps.StatusCode)
		return
	}

	result, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", result)
}
