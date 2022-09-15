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
	tanksSprite := sprites["tank_blue"]
	println(tanksSprite.Frame().String())

	cfg := pixelgl.WindowConfig{
		Title:     "Pixel Rocks!",
		Bounds:    pixel.R(0, 0, 480, 360),
		VSync:     true,
		Resizable: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	tank := MakeTank(win, 300, radians(200))

	// make walls
	walls_ := make([]Entity, 0, 10)
	for i := 0; i < 10; i++ {
		walls_ = append(walls_, NewWall(pixel.V(float64(i*56), 0)))
	}
	walls := MakeGroup(walls_...)

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if win.JustPressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Greenyellow)

		tank.Update(dt, win)
		tank.Draw(win)

		walls.Update(dt, win)
		walls.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
