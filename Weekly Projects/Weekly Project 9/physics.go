package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsBody struct {
	*Transform
	Vel rl.Vector2
}

func (b PhysicsBody) GetRectangle() rl.Rectangle {
	return rl.NewRectangle(b.Pos.X-b.Scale.X/2, b.Pos.Y-b.Scale.Y/2, b.Scale.X, b.Scale.Y)
}

func (b PhysicsBody) GetPlayerRectangle() rl.Rectangle {
	return rl.NewRectangle(64+b.Pos.X-b.Scale.X/2, b.Pos.Y-b.Scale.Y/2, b.Scale.X/2, b.Scale.Y)
}

func (b *PhysicsBody) GetPlayerPunchHitBox() rl.Rectangle {
	if b.Flip == 1 {
		return rl.NewRectangle(180+b.Pos.X-b.Scale.X/2, 32+b.Pos.Y-b.Scale.Y/2, b.Scale.X/4, b.Scale.Y/2)
	} else {
		return rl.NewRectangle(12+b.Pos.X-b.Scale.X/2, 32+b.Pos.Y-b.Scale.Y/2, b.Scale.X/4, b.Scale.Y/2)
	}
}

func (b *PhysicsBody) GetPlayerKickHitBox() rl.Rectangle {
	if b.Flip == 1 {
		return rl.NewRectangle(180+b.Pos.X-b.Scale.X/2, 64+b.Pos.Y-b.Scale.Y/2, b.Scale.X/4, b.Scale.Y/2)
	} else {
		return rl.NewRectangle(12+b.Pos.X-b.Scale.X/2, 64+b.Pos.Y-b.Scale.Y/2, b.Scale.X/4, b.Scale.Y/2)
	}
}

func (b *PhysicsBody) GetPlayerBlockHitBox() rl.Rectangle {
	if b.Flip == 1 {
		return rl.NewRectangle(174+b.Pos.X-b.Scale.X/2, b.Pos.Y-b.Scale.Y/2, -32+b.Scale.X/4, b.Scale.Y)
	} else {
		return rl.NewRectangle(50+b.Pos.X-b.Scale.X/2, b.Pos.Y-b.Scale.Y/2, -32+b.Scale.X/4, b.Scale.Y)
	}
}

func (b *PhysicsBody) ApplyGravity(g rl.Vector2) {
	b.Vel = rl.Vector2Add(b.Vel, rl.Vector2Scale(g, rl.GetFrameTime()))
}

func (b *PhysicsBody) ApplyVel() {
	b.Pos = rl.Vector2Add(b.Pos, rl.Vector2Scale(b.Vel, rl.GetFrameTime()))
}
