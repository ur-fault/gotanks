package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"math"
)

type Tank struct {
	entity
	speed          float64
	rotSpeed       float64
	barrel         *pixel.Sprite
	barrelRotation float64 // in radians, relative to the tank
	debugPoint     pixel.Vec
	velocity       float64
	acceleration   float64
}

func MakeTank(speed float64, rotSpeed float64, acceleration float64) Tank {
	tank := Tank{}
	tank.sprite = sprites["tankBody_green_outline"]
	tank.position = pixel.ZV
	tank.scale = 1
	tank.speed = speed
	tank.rotSpeed = rotSpeed
	tank.barrel = sprites["tankGreen_barrel2_outline"]
	tank.acceleration = acceleration
	return tank
}

func (t *Tank) Update(dt float64, win *pixelgl.Window) {
	targetVelocity := 0.0
	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		t.rotation += t.rotSpeed * dt
	}
	if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
		t.rotation -= t.rotSpeed * dt
	}
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
		targetVelocity = t.speed
	}

	if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
		targetVelocity = -t.speed
	}

	if t.velocity < targetVelocity {
		t.velocity = math.Min(t.velocity+t.acceleration*dt, targetVelocity)
	}
	if t.velocity > targetVelocity {
		t.velocity = math.Max(t.velocity-t.acceleration*dt, targetVelocity)
	}

	t.position.X += t.velocity * dt * math.Cos(t.Facing())
	t.position.Y += t.velocity * dt * math.Sin(t.Facing())
}

func (t *Tank) Draw(cam *Camera) {
	// call default Draw
	t.entity.Draw(cam)
	win := cam.Window()

	barrelMatrix := pixel.IM.
		Rotated(pixel.ZV, radians(180)+t.barrelRotation).
		Moved(pixel.V(0, 7)).
		Chained(t.Matrix())

	t.barrel.Draw(win, barrelMatrix)

	im := imdraw.New(nil)
	im.Color = colornames.Aliceblue
	im.Push(cam.ViewPortToWorld(win.MousePosition().Sub(win.Bounds().Center())))
	im.Circle(cam.ToUnits(2), 0)
	im.Draw(win)
}

func (t *Tank) Rotate(angle float64) {
	t.rotation += angle
}

func (t *Tank) Facing() float64 {
	return t.rotation - radians(90)
}
