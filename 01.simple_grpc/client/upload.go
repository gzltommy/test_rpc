package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"simple_grpc/proto/rpc"
	"time"
)

func Upload(grpcConn *grpc.ClientConn) {
	//通过grpc连接创建一个客户端实例对象
	client := rpc.NewUploadClient(grpcConn)

	//设置ctx超时（根据情况设定）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//和简单rpc不同，此时获得的不是res，而是一个client的对象，通过这个连接对象去发送数据
	uploadClient, err := client.Upload(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	var offset int64
	var size int64
	size = 4 * 1024

	// 循环处理数据，当大于 64kb 退出
	for {
		err := uploadClient.Send(&rpc.UploadReq{
			Path:   "../test_grpc.txt",
			Offset: offset,
			Size:   size,
			Data:   nil,
		})
		if err != nil {
			log.Fatalln(err)
		}
		offset += size
		//发送超过64KB，调用CloseAndRecv方法接收response
		if offset >= 64*1024 {
			res, err := uploadClient.CloseAndRecv()
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("upload over~, response is ", res.Msg)
			break
		}
	}
}
