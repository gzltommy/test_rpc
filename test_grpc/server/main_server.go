// main_server.go
package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"test_grpc/proto/rpc"
)

func main() {
	lis, err := net.Listen("tcp", ":6012")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//构建一个新的服务端对象
	s := grpc.NewServer()

	//向这个服务端对象注册服务
	rpc.RegisterLoginServer(s, &LoginServer{})
	rpc.RegisterUploadServer(s, &UploadServer{})
	rpc.RegisterDownloadServer(s, &DownloadServer{})

	//注册服务端反射服务
	reflection.Register(s)

	//启动服务
	s.Serve(lis)

	//可配合ctx实现服务端的动态终止
	//s.Stop()
}