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
		return pixel.IM.Moved(c.parent.Position().Scaled(-1)).Moved(c.window.Bounds().Center()).Rotated(c.window.Bounds().Center(), -c.parent.Rotation()).Scaled(c.window.Bounds().Center(), c.PixelsPerUnit())
	} else {
		return pixel.IM.Moved(c.parent.Position().Scaled(-1)).Moved(c.window.Bounds().Center()).Scaled(c.window.Bounds().Center(), c.PixelsPerUnit())
	}
}

func (c *Camera) RevMatrix() pixel.Matrix {
	if c.rotate {
		return pixel.IM.Moved(c.parent.Position()).Rotated(c.parent.Position(), c.parent.Rotation()).Scaled(c.parent.Position(), c.UnitsPerPixel())
	} else {
		return pixel.IM.Moved(c.parent.Position()).Scaled(c.parent.Position(), c.UnitsPerPixel())
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

// Apply applies the camera's matrix to the window. Should be called before drawing or changing space (like drawing UI)
func (c *Camera) Apply() {
	c.window.SetMatrix(c.Matrix())
}

func (c *Camera) ViewPortToWorld(viewport pixel.Vec) pixel.Vec {
	return c.RevMatrix().Project(viewport)
}

func (c *Camera) WorldToViewPort(world pixel.Vec) pixel.Vec {
	return c.Matrix().Project(world)
}

func (c *Camera) UnitsPerPixel() float64 {
	return 1 / (c.Window().Bounds().Size().Y * c.zoom)
}

func (c *Camera) PixelsPerUnit() float64 {
	return c.zoom * c.Window().Bounds().Size().Y
}

func (c *Camera) ToUnits(pixels float64) float64 {
	return pixels * c.UnitsPerPixel()
}

func (c *Camera) ToPixels(units float64) float64 {
	return units * c.PixelsPerUnit()
}
