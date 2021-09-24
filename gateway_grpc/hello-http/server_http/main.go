package main

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	gw "gateway_grpc/proto/hello_http"
)

const (
	Endpoint   = "127.0.0.1:50052" // grpc 服务地址
	HttpListen = "0.0.0.0:8080"    // http 服务的监听地址
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// HTTP 转 grpc
	err := gw.RegisterHelloHandlerFromEndpoint(ctx, mux, Endpoint, opts)
	if err != nil {
		grpclog.Fatalf("Register handler err:%v\n", err)
	}

	fmt.Println("HTTP Listen on 8080")

	http.ListenAndServe(HttpListen, mux)
}
