package client

import (
	"context"
	"io"

	"github.com/Lunarisnia/stream-demo/chatservice"
	"github.com/veandco/go-sdl2/sdl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func drawPixel(s *sdl.Surface, x int, y int, color *sdl.Color) {
	rect := sdl.Rect{X: int32(x), Y: int32(y), W: 1, H: 1}
	// colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(s.Format, color.R, color.G, color.B, color.A)
	s.FillRect(&rect, pixel)
}

func Connect() (chatservice.ChatServiceClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("127.0.0.1:3009", opts...)
	if err != nil {
		return nil, err
	}

	cs := chatservice.NewChatServiceClient(conn)

	return cs, err
}

func Render(cs chatservice.ChatServiceClient, window *sdl.Window, surface *sdl.Surface, callback func()) error {
	stream, err := cs.Render(context.Background(), &chatservice.RenderOption{
		Width:  100,
		Height: 100,
	})
	if err != nil {
		return err
	}
	for {
		pixel, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		drawPixel(surface, int(pixel.X), int(pixel.Y), &sdl.Color{
			R: uint8(pixel.Color.R),
			G: uint8(pixel.Color.G),
			B: uint8(pixel.Color.B),
			A: 255,
		})
		window.UpdateSurface()
	}
	callback()

	return nil
}
