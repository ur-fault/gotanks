package main

import "github.com/faiface/pixel"

type Wall struct {
	entity
}

var wallSize float64 = 2

func MakeWall(pos pixel.Vec) Wall {
	wall := Wall{}
	wall.sprite = sprites["crateMetal"]
	wall.scale = 1
	wall.position = pos
	return wall
}

func NewWall(pos pixel.Vec) *Wall {
	wall := MakeWall(pos)
	wall.SetWidth(56)
	return &wall
}

func FillWalls(pos pixel.Vec, width, height int, spacingX, spacingY float64) []Entity {
	walls := make([]Entity, 0, width*height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			wall := NewWall(pixel.V(pos.X+float64(i)*(wallSize*spacingX), pos.Y+float64(j)*(wallSize*spacingY)))
			wall.SetWidth(wallSize)
			walls = append(walls, wall)
		}
	}
	return walls
}
