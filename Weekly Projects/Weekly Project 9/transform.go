package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Transform struct {
	Pos   rl.Vector2
	Flip  int
	Scale rl.Vector2
}

func NewTransform(newPos rl.Vector2) *Transform {
	return &Transform{Pos: newPos, Flip: 1, Scale: rl.NewVector2(1, 1)}
}
