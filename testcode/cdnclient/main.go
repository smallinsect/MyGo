package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
	// 创建OSSClient实例。
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", "LTAI4FhVu9miaTN5yeFXpztk", "CQQ6CCgFh7Iz76PiG0XXihmqPfQVMa")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("panzerwy4a")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile("Maintenance.json", "D:/Maintenance.json")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
