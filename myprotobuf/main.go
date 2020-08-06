package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"xuqiulin.com/mygo/myprotobuf/example"
)

func main() {
	fmt.Println("小昆虫")

	msgTest := &example.Person{
		Name: proto.String("小白菜"),
		Age:  proto.Int32(666),
		From: proto.String("中国"),
	}

	// 序列化
	msgDataEncoding, err := proto.Marshal(msgTest)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgEntity := example.Person{}
	// 反序列化
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("姓名：", msgEntity.GetName())
	fmt.Println("年龄：", msgEntity.GetAge())
	fmt.Println("国籍：", msgEntity.GetFrom())
}
