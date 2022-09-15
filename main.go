package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	_ "image/jpeg"
	_ "image/png"
	"time"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Go, Tanks",
		Bounds:    pixel.R(0, 0, 500, 500),
		VSync:     true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	tank := MakeTank(20, radians(200), 55)
	tank.SetWidth(3)
	println(tank.Scale())

	cam := MakeCamera(&tank.entity, win)
	cam.SetZoom(1.0 / 30)
	println(cam.PixelsPerUnit())

	walls := MakeGroup(FillWalls(pixel.V(0, 0), 10, 1, 1, 1)...)
	walls.AddEntities(FillWalls(pixel.V(0, 0), 1, 10, 1, 1)...)

	start := time.Now()
	last := start
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if win.JustPressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Firebrick)

		walls.Update(dt, win)
		tank.Update(dt, win)

		cam.Apply()

		walls.Draw(&cam)
		tank.Draw(&cam)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
