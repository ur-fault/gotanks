package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Entity interface {
	// GetPosition returns the position of the entity
	GetPosition() pixel.Vec
	// GetSprite returns the sprite of the entity
	GetSprite() *pixel.Sprite
	// GetRotation returns the rotation in radians of the entity
	GetRotation() float64
	// GetScale returns the scale of the entity
	GetScale() float64
	// SetPosition sets the position of the entity
	SetPosition(position pixel.Vec)
	// SetRotation sets the rotation in radians of the entity
	SetRotation(rotation float64)
	// SetScale sets the scale of the entity
	SetScale(scale float64)
	// Update updates the entity
	Update(dt float64, win *pixelgl.Window)
	// Draw draws the entity
	Draw(win *pixelgl.Window)
	// GetSize returns the size of the entity
	GetSize() pixel.Vec
	// GetBounds returns the bounds of the entity
	GetBounds() pixel.Rect
	// SetWidth sets the width of the entity
	SetWidth(width float64)
}

type entity struct {
	position pixel.Vec
	sprite   *pixel.Sprite
	rotation float64
	scale    float64
}

func (e *entity) GetPosition() pixel.Vec {
	return e.position
}

func (e *entity) GetSprite() *pixel.Sprite {
	return e.sprite
}

func (e *entity) GetRotation() float64 {
	return e.rotation
}

func (e *entity) GetScale() float64 {
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

func (e *entity) Draw(win *pixelgl.Window) {
	e.sprite.Draw(win, pixel.IM.Moved(e.position).Rotated(e.position, e.rotation).Scaled(e.position, e.scale))
}

func (e *entity) GetSize() pixel.Vec {
	return e.sprite.Frame().Size()
}

func (e *entity) GetBounds() pixel.Rect {
	return e.sprite.Frame()
}

func (e *entity) SetWidth(width float64) {
	e.scale = width / e.GetSize().X
}
