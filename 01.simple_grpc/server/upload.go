// upload.go

package main

import (
	"io"
	"log"
	"simple_grpc/proto/rpc"
)

type UploadServer struct{}

func (*UploadServer) Upload(uploadServer rpc.Upload_UploadServer) error {
	for {
		//循环接受客户端传的流数据
		recv, err := uploadServer.Recv()

		//检测到EOF（客户端调用close）
		if err == io.EOF {
			//发送res
			err := uploadServer.SendAndClose(&rpc.UploadRes{Msg: "finish"})
			if err != nil {
				return err
			}
			return nil
		} else if err != nil {
			return err
		}
		log.Printf("get a upload data package~ offset:%v, size:%v\n", recv.Offset, recv.Size)
	}
}
