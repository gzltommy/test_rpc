// main_client.go
package main

import (
	"google.golang.org/grpc"
	"log"
)

func main() {
	//创立 grpc 连接
	grpcConn, err := grpc.Dial("127.0.0.1:6012", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	// 登录
	Login(grpcConn)

	// 上传
	Upload(grpcConn)

	// 下载
	Download(grpcConn)
}
