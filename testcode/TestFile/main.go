// main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type User struct {
	Name string
	Age  int8
}

func testMarshal() []byte {
	user := User{
		Name: "小昆虫",
		Age:  20,
	}
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func testUnmarshal(data []byte) {
	var user User
	err := json.Unmarshal(data, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)
}

func testRead() []byte {
	fp, err := os.OpenFile("./data.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data[:n]))
	return data[:n]
}

func testWrite(data []byte) {
	os.Remove("data.json")

	fp, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var data []byte
	data = testMarshal()
	fmt.Println(string(data))
	testWrite(data)
	data = testRead()
	testUnmarshal(data)
}

func Data1() {
	b, err := ioutil.ReadFile("debug-log.202011170800")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(b)
	contents := strings.Split(content, "\n")

	result := make([]string, 0)
	i := 0
	lenght := len(contents)
	for i < lenght {
		if strings.Contains(contents[i], "OnClose") {
			result = append(result, contents[i])
		}
		i++
	}
	fl, err := os.OpenFile("1", os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	i = 0
	for i < len(result) {
		fl.Write([]byte(result[i]))
		fl.Write([]byte("\n"))
		i++
	}
	fmt.Println(len(contents))
}

func Data2() {
	b, err := ioutil.ReadFile("1")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(b)
	contents := strings.Split(content, "\n")

	result := make([]string, 0)
	i := 0
	lenght := len(contents)
	for i < lenght {
		if strings.Contains(contents[i], "172.20.101.40:33358") {
			result = append(result, contents[i])
		}
		i++
	}
	fl, err := os.OpenFile("2", os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	i = 0
	for i < len(result) {
		fl.Write([]byte(result[i]))
		fl.Write([]byte("\n"))
		i++
	}
	fmt.Println(len(contents))
}
