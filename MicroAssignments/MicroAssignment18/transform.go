package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Transform struct {
	Pos   rl.Vector2
	Flip  int
	Scale int
	Angle float32
}

func NewTransform(newPos rl.Vector2) *Transform {
	return &Transform{Pos: newPos, Flip: 1, Scale: 1, Angle: 0}
}
