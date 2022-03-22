package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)



func main() {
	// 创建OSSClient实例。
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "xxxxxxxxxx", "xxxxxxxxxxxxxxxxx")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("mimall1700")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 读取本地文件。
	fd, err := os.Open("C:\\Users\\A_\\Pictures\\jiang.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer fd.Close()

	// 上传文件流。
	err = bucket.PutObject("2020/xxxx/beego.jpg", fd)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
