package main

import (
	"context"
	"fmt"
	"io"

	"github.com/Lunarisnia/stream-demo/chatservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("127.0.0.1:3009", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cs := chatservice.NewChatServiceClient(conn)
	response, err := cs.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server:", response.Message)

	stream, err := cs.Render(context.Background(), &chatservice.RenderOption{
		Width:  100,
		Height: 100,
	})
	for {
		pixel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("RECEIVED PIXEL: ", pixel)
	}
}
