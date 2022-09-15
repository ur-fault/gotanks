package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type Tank struct {
	entity
	speed          float64
	rotSpeed       float64
	barrel         *pixel.Sprite
	barrelRotation float64 // in radians, relative to the tank
	debugPoint     pixel.Vec
}

func MakeTank(win *pixelgl.Window, speed float64, rotSpeed float64) Tank {
	tank := Tank{}
	tank.sprite = sprites["tankBody_green_outline"]
	tank.position = pixel.ZV
	tank.scale = 1
	tank.speed = speed
	tank.rotSpeed = rotSpeed
	tank.barrel = sprites["tankGreen_barrel2_outline"]
	return tank
}

func (t *Tank) Update(dt float64, win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		t.rotation += t.rotSpeed * dt
	}
	if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
		t.rotation -= t.rotSpeed * dt
	}
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
		t.position.X += t.speed * dt * math.Cos(t.Facing())
		t.position.Y += t.speed * dt * math.Sin(t.Facing())
	}

	if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
		t.position.X -= t.speed * dt * math.Cos(t.Facing())
		t.position.Y -= t.speed * dt * math.Sin(t.Facing())
	}
}

func (t *Tank) Draw(cam *Camera) {
	win := cam.Window()
	t.sprite.Draw(win, t.Matrix())

	barrelMatrix := pixel.IM.
		Rotated(pixel.ZV, radians(180)+t.barrelRotation).
		Moved(pixel.V(0, 7)).
		Chained(t.Matrix())

	t.barrel.Draw(win, barrelMatrix)

	im := imdraw.New(win.Canvas())
	im.Color = pixel.RGB(1, 0, 0)
	im.SetMatrix(cam.RevMatrix())
	println(win.MousePosition().Sub(win.Bounds().Center()).String())
	im.Push(win.MousePosition().Sub(win.Bounds().Center()))
	im.Circle(5, 0)
	im.Draw(win)
}

func (t *Tank) Rotate(angle float64) {
	t.rotation += angle
}

func (t *Tank) Facing() float64 {
	return t.rotation - radians(90)
}
