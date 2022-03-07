package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	hello_grpc "wmh/protoc/proto"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) { //context 上下文
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "服务端返回的内容"}, nil
}

func main() {
	// 取出服务

	// 挂载方法

	// 注册服务

	// 创建监听
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
