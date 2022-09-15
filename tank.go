package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"math"
)

type Tank struct {
	entity
	speed    float64
	rotSpeed float64
}

func MakeTank(win *pixelgl.Window, speed float64, rotSpeed float64) Tank {
	tank := Tank{}
	tank.sprite = sprites["tankBody_green_outline"]
	tank.position = win.Bounds().Center()
	tank.scale = 1
	tank.speed = speed
	tank.rotSpeed = rotSpeed
	return tank
}

func (t *Tank) Update(dt float64, win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyA) {
		t.rotation += t.rotSpeed * dt
	}
	if win.Pressed(pixelgl.KeyD) {
		t.rotation -= t.rotSpeed * dt
	}
	if win.Pressed(pixelgl.KeyW) {
		t.position.X += t.speed * dt * math.Cos(t.correctDir())
		t.position.Y += t.speed * dt * math.Sin(t.correctDir())
	}

	if win.Pressed(pixelgl.KeyS) {
		t.position.X -= t.speed * dt * math.Cos(t.correctDir())
		t.position.Y -= t.speed * dt * math.Sin(t.correctDir())
	}
}

func (t *Tank) Draw(win *pixelgl.Window) {
	t.sprite.Draw(win, pixel.IM.Moved(t.position).Rotated(t.position, t.rotation).Scaled(t.position, t.scale))
}

func (t *Tank) rotate(angle float64) {
	t.rotation += angle
}

func (t *Tank) correctDir() float64 {
	return t.rotation - radians(90)
}
