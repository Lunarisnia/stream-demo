package server

import (
	"context"
	"math/rand"
	"time"

	"github.com/Lunarisnia/stream-demo/chatservice"
)

type chatServiceServer struct {
	chatservice.UnimplementedChatServiceServer
}

func NewChatServiceServer() chatservice.ChatServiceServer {
	return &chatServiceServer{}
}

func (c *chatServiceServer) Ping(ctx context.Context, in *chatservice.PingMessage) (*chatservice.PingResponse, error) {
	return &chatservice.PingResponse{
		Message: "Pong",
	}, nil
}

func (c *chatServiceServer) Render(renderOption *chatservice.RenderOption, stream chatservice.ChatService_RenderServer) error {
	for i := range 100 {
		err := stream.Send(&chatservice.Pixel{
			X: int32(i),
			Y: int32(i),
			Color: &chatservice.Color{
				R: int32(rand.Intn(256)),
				G: int32(rand.Intn(256)),
				B: int32(rand.Intn(256)),
			},
		})
		if err != nil {
			return err
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil
}
