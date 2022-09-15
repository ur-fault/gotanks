package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type Camera interface {
	// GetPosition returns the position of the camera
	GetPosition() pixel.Vec
	// GetRotation returns the rotation in radians of the camera
	GetRotation() float64
	// GetZoom returns the scale of the camera
	GetZoom() float64
	// SetPosition sets the position of the camera
	SetPosition(position pixel.Vec)
	// SetRotation sets the rotation in radians of the camera
	SetRotation(rotation float64)
	// SetZoom sets the scale of the camera
	SetZoom(zoom float64)
	// GetWindow returns the window of the camera
	GetWindow() *pixelgl.Window
	// GetWindowSize returns the size of the window
	GetWindowSize() pixel.Vec
	// GetMatrix returns the matrix of the camera
	GetMatrix() pixel.Matrix
	// GetRatio returns the ratio of the camera
	GetRatio() float64
	// ToViewpoint converts a point from the world to the camera's viewpoint
	ToViewpoint(point pixel.Vec) pixel.Vec
	// ToWorld converts a point from the camera's viewpoint to the world
	ToWorld(point pixel.Vec) pixel.Vec
}
