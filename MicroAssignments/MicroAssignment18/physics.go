package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsBody struct {
	*Transform
	Vel  rl.Vector2
	Size rl.Vector2
}

func (b *PhysicsBody) ApplyGravity(g rl.Vector2) {
	b.Vel = rl.Vector2Add(b.Vel, rl.Vector2Scale(g, rl.GetFrameTime()))
}

func (b *PhysicsBody) ApplyVel() {
	b.Pos = rl.Vector2Add(b.Pos, rl.Vector2Scale(b.Vel, rl.GetFrameTime()))
}
