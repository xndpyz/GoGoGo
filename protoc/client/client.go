package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "wmh/protoc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println("client 错误：", err)
	}
	defer conn.Close()
	client := hello_grpc.NewHelloGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc.Req{Message: "我从客户端来"})
	fmt.Println(req.GetMessage())
}
