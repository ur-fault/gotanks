package main

import "github.com/faiface/pixel"

type Wall struct {
	entity
}

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
