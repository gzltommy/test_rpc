// download.go

package main

import (
	"test_grpc/proto/rpc"
)

type DownloadServer struct{}

func (*DownloadServer) Download(req *rpc.DownloadReq, downloadServer rpc.Download_DownloadServer) error {
	offset := req.Offset
	//循环发送数据
	for {
		err := downloadServer.Send(&rpc.DownloadRes{
			Offset: offset,
			Size:   4 * 1024,
			Data:   nil,
		})
		if err != nil {
			return err
		}
		offset += 4 * 1024
		if offset >= req.Offset+req.Size {
			break
		}
	}
	return nil
}
