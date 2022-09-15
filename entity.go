package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Entity interface {
	// Position returns the position of the entity
	Position() pixel.Vec
	// Sprite returns the sprite of the entity
	Sprite() *pixel.Sprite
	// Rotation returns the rotation in radians of the entity
	Rotation() float64
	// Scale returns the scale of the entity
	Scale() float64
	// SetPosition sets the position of the entity
	SetPosition(position pixel.Vec)
	// SetRotation sets the rotation in radians of the entity
	SetRotation(rotation float64)
	// SetScale sets the scale of the entity
	SetScale(scale float64)
	// Update updates the entity
	Update(dt float64, win *pixelgl.Window)
	// Draw draws the entity
	Draw(cam *Camera)
	// Size returns the size of the entity
	Size() pixel.Vec
	// Bounds returns the bounds of the entity
	Bounds() pixel.Rect
	// SetWidth sets the width of the entity
	SetWidth(width float64)
	// Matrix returns the matrix of the entity
	Matrix() pixel.Matrix
}

type entity struct {
	position pixel.Vec
	sprite   *pixel.Sprite
	rotation float64
	scale    float64
}

func (e *entity) Position() pixel.Vec {
	return e.position
}

func (e *entity) Sprite() *pixel.Sprite {
	return e.sprite
}

func (e *entity) Rotation() float64 {
	return e.rotation
}

func (e *entity) Scale() float64 {
	return e.scale
}

func (e *entity) SetPosition(position pixel.Vec) {
	e.position = position
}

func (e *entity) SetRotation(rotation float64) {
	e.rotation = rotation
}

func (e *entity) SetScale(scale float64) {
	e.scale = scale
}

func (e *entity) Update(dt float64, win *pixelgl.Window) {
	// TODO
}

func (e *entity) Draw(cam *Camera) {
	e.sprite.Draw(cam.Window(), e.Matrix())
}

func (e *entity) Size() pixel.Vec {
	return e.sprite.Frame().Size()
}

func (e *entity) Bounds() pixel.Rect {
	return e.sprite.Frame()
}

func (e *entity) SetWidth(width float64) {
	e.SetScale(width / e.Size().X)
}

func (e *entity) Matrix() pixel.Matrix {
	return pixel.IM.Moved(e.position).Rotated(e.position, e.rotation).Scaled(e.position, e.scale)
}
