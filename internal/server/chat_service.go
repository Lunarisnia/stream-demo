package server

import (
	"context"

	"github.com/Lunarisnia/stream-demo/chatservice"
	"github.com/lunarisnia/yacg"
	"github.com/lunarisnia/yacg/color"
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
	yacg.PathTrace(10, 400, func(x int, y int, c *color.RGB) {
		stream.Send(&chatservice.Pixel{
			X: int32(x),
			Y: int32(y),
			Color: &chatservice.Color{
				R: int32(c.Red),
				G: int32(c.Green),
				B: int32(c.Blue),
			},
		})
	})
	return nil
}
