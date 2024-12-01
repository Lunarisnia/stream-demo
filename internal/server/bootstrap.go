package server

import (
	"fmt"
	"net"

	"github.com/Lunarisnia/stream-demo/chatservice"
	"google.golang.org/grpc"
)

func Run() error {
	lis, err := net.Listen("tcp", ":3009")
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	chatservice.RegisterChatServiceServer(grpcServer, NewChatServiceServer())

	fmt.Println("Starting Server at 3009")
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
