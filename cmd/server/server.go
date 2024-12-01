package main

import (
	"fmt"
	"net"

	"github.com/Lunarisnia/stream-demo/chatservice"
	"github.com/Lunarisnia/stream-demo/internal/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3009")
	if err != nil {
		panic(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	chatservice.RegisterChatServiceServer(grpcServer, server.NewChatServiceServer())

	fmt.Println("Starting Server at 3009")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
