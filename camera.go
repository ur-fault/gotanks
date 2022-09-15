package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera struct {
	parent Entity
	rotate bool
	zoom   float64
	window *pixelgl.Window
}

func MakeCamera(parent Entity, window *pixelgl.Window) Camera {
	return Camera{
		parent: parent,
		rotate: false,
		zoom:   1,
		window: window,
	}
}

func NewCamera(parent Entity, window *pixelgl.Window) *Camera {
	cam := MakeCamera(parent, window)
	return &cam
}

func (c *Camera) Parent() Entity {
	return c.parent
}

func (c *Camera) SetParent(parent Entity) {
	c.parent = parent
}

func (c *Camera) Rotate() bool {
	return c.rotate
}

func (c *Camera) SetRotate(rotate bool) {
	c.rotate = rotate
}

func (c *Camera) Matrix() pixel.Matrix {
	if c.rotate {
		return pixel.IM.Moved(c.parent.Position().Scaled(-1)).Moved(c.window.Bounds().Center()).Rotated(c.window.Bounds().Center(), -c.parent.Rotation())
	} else {
		return pixel.IM.Moved(c.parent.Position().Scaled(-1)).Moved(c.window.Bounds().Center())
	}
}

func (c *Camera) RevMatrix() pixel.Matrix {
	if c.rotate {
		return pixel.IM.Moved(c.parent.Position()).Rotated(c.parent.Position(), c.parent.Rotation())
	} else {
		return pixel.IM.Moved(c.parent.Position())
	}
}

func (c *Camera) Zoom() float64 {
	return c.zoom
}

func (c *Camera) SetZoom(zoom float64) {
	c.zoom = zoom
}

func (c *Camera) Window() *pixelgl.Window {
	return c.window
}

func (c *Camera) SetWindow(window *pixelgl.Window) {
	c.window = window
}

func (c *Camera) Apply() {
	c.window.SetMatrix(c.Matrix())
}
