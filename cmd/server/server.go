package main

import (
	"github.com/Lunarisnia/stream-demo/internal/server"
	"github.com/gin-gonic/gin"
	"github.com/lunarisnia/yacg"
	"github.com/lunarisnia/yacg/color"
)

type Color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type Pixel struct {
	X     int   `json:"x"`
	Y     int   `json:"y"`
	Color Color `json:"color"`
}

func main() {
	go server.Run()
	r := gin.Default()
	r.GET("/render", func(c *gin.Context) {
		pixels := make([]Pixel, 0)
		yacg.PathTrace(500, 400, func(x int, y int, c *color.RGB) {
			pixels = append(pixels, Pixel{
				X: x,
				Y: y,
				Color: Color{
					R: c.Red,
					G: c.Green,
					B: c.Blue,
				},
			})
		})
		c.JSON(200, gin.H{
			"data": pixels,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
