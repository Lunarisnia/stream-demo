package render

import (
	"fmt"
	"math"

	"github.com/Lunarisnia/stream-demo/internal/client"
	"github.com/veandco/go-sdl2/sdl"
)

func StartRenderer() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	aspectRatio := 16.0 / 9.0
	screenWidth := 400
	screenHeight := int(math.Round(float64(screenWidth) / aspectRatio))
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(screenWidth), int32(screenHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	chatService, err := client.Connect()
	if err != nil {
		panic(err)
	}

	rendering := false
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				keyEvent := event.(*sdl.KeyboardEvent)
				if keyEvent.Keysym.Scancode == sdl.SCANCODE_S && keyEvent.Type == sdl.KEYDOWN {
					surface.FillRect(nil, 0)
					window.UpdateSurface()
				}
				if keyEvent.Keysym.Scancode == sdl.SCANCODE_R && keyEvent.Type == sdl.KEYDOWN {
					if !rendering {
						surface.FillRect(nil, 0)
						window.UpdateSurface()
						fmt.Println("RAASKDJASDKLj")
						go func() {
							err := client.Render(chatService, window, surface, func() {
								rendering = false
							})
							if err != nil {
								fmt.Println(err)
							}
						}()
						rendering = true
					}
				}
				if keyEvent.Keysym.Scancode == sdl.SCANCODE_H && keyEvent.Type == sdl.KEYDOWN {
					if !rendering {
						surface.FillRect(nil, 0)
						window.UpdateSurface()
						go client.Render(chatService, window, surface, func() {
							rendering = false
						})
						rendering = true
					}
				}
			}
		}

		sdl.Delay(33)
	}
}
