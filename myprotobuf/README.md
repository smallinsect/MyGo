# protocolbuffers

https://github.com/protocolbuffers/protobuf

```
git clone https://github.com/protocolbuffers/protobuf
```

https://github.com/protocolbuffers/protobuf-go



https://github.com/protocolbuffers/protobuf/tree/master/src



# 前锋教育笔记

Github地址

- https://github.com/golang/protobuf

选择适合自己系统的protobuf编译器

- https://github.com/protocolbuffers/protobuf/releases

```
生成Go文件
protoc --go_out=. test.proto
生成CPP文件
protoc --cpp_out=. test.proto
```



安装Go下载源码

```
go get -u -v github.com/golang/protobuf/protoc-gen-go
```

- https://github.com/golang/protobuf



# protoc使用

```
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
```

# 语法

```
syntax = "proto3";
option go_package = ".;example";
package example;

// message 消息类型
message User {
    string name = 1;
    int32 age   = 2;
    string from = 3;
    int64 price = 4;
}
```

