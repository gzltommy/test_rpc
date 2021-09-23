// main_client.go
package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"test_grpc/proto/rpc"
	"time"
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

func Login(grpcConn *grpc.ClientConn) {
	//通过grpc连接创建一个客户端实例对象
	client := rpc.NewLoginClient(grpcConn)

	//设置ctx超时（根据情况设定）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//通过client客户端对象，调用Login函数
	res, err := client.Login(ctx, &rpc.LoginReq{
		Username: "root",
		Password: "123456",
	})
	if err != nil {
		log.Fatalln(err)
	}

	//输出登陆结果
	log.Println("the login answer is", res.Msg)
}


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

func Download(grpcConn *grpc.ClientConn) {
	//通过grpc连接创建一个客户端实例对象
	client := rpc.NewDownloadClient(grpcConn)

	//设置ctx超时（根据情况设定）
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//和简单rpc不同，此时获得的不是res，而是一个client的对象，通过这个连接对象去读取数据
	downloadClient, err := client.Download(ctx, &rpc.DownloadReq{
		Path:   "../test_grpc.txt",
		Offset: 0,
		Size:   64 * 1024,
	})
	if err != nil {
		log.Fatalln(err)
	}

	//循环处理数据，当监测到读取完成后退出
	for {
		res, err := downloadClient.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("get a date package~ offset:%v, size:%v\n", res.Offset, res.Size)
		if res.Size+res.Offset >= 64*1024 {
			break
		}
	}

	log.Println("download over~")
}