package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsBody struct {
	position rl.Vector2
	velocity rl.Vector2
	radius   float32 // radius for collision purposes, not drawing purposes
	angle    float32
	texture  rl.Texture2D
	collide  bool
}

// create a new physics body
func newPhysicsBody(pos rl.Vector2, vel rl.Vector2, rad float32, tex rl.Texture2D, col bool) PhysicsBody {
	return PhysicsBody{position: pos, velocity: vel, radius: rad, angle: 0, texture: tex, collide: col}
}

// see week 5 monday slides for below
func DrawBody(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)
	rl.DrawTexturePro(texture, sourceRect, destRect, origin, angle, color)
}
