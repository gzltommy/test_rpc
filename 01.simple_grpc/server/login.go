// login.go

package main

import (
	"context"
	"simple_grpc/proto/rpc"
)

type LoginServer struct{}

//判断用户名，密码是否为root,123456，验证正确即返回
func (*LoginServer) Login(ctx context.Context, req *rpc.LoginReq) (*rpc.LoginRes, error) {
	//为降低复杂度，此处不对 ctx 进行处理
	if req.Username == "root" && req.Password == "123456" {
		return &rpc.LoginRes{Msg: "true"}, nil
	} else {
		return &rpc.LoginRes{Msg: "false"}, nil
	}
}
