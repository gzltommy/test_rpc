package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"simple_grpc/proto/rpc"
	"time"
)

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
