package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	PhysicsBody
	cargo int
	speed int
}

// creates the player
func createPlayer(playerTexture rl.Texture2D) Player {
	pb := newPhysicsBody(rl.NewVector2(640, 360), rl.NewVector2(0, 0), 16, playerTexture, false)
	return Player{cargo: 0, speed: 300, PhysicsBody: pb}
}

func (p *Player) playerMove() {
	// player movement - normalized diagonal movement
	tempVelocity := rl.Vector2Zero()
	if rl.IsKeyDown(rl.KeyW) {
		tempVelocity.Y = -1 * float32(p.speed)
	}
	if rl.IsKeyDown(rl.KeyA) {
		tempVelocity.X = -1 * float32(p.speed)
	}
	if rl.IsKeyDown(rl.KeyS) {
		tempVelocity.Y = float32(p.speed)
	}
	if rl.IsKeyDown(rl.KeyD) {
		tempVelocity.X = float32(p.speed)
	}
	if tempVelocity.X == 0 || tempVelocity.Y == 0 {
		p.position = rl.Vector2Add(p.position, rl.Vector2Scale(tempVelocity, rl.GetFrameTime()))
	} else {
		x := float64(tempVelocity.X)
		y := float64(tempVelocity.Y)
		v := math.Sqrt(x*x + y*y)
		tempVelocity.X = (tempVelocity.X / float32(v)) * 300
		tempVelocity.Y = (tempVelocity.Y / float32(v)) * 300
		p.position = rl.Vector2Add(p.position, rl.Vector2Scale(tempVelocity, rl.GetFrameTime()))
	}

	// set screen boundary
	if p.position.X > 1280 {
		p.position.X = 1280
	} else if p.position.X < 0 {
		p.position.X = 0
	} else if p.position.Y > 720 {
		p.position.Y = 720
	} else if p.position.Y < 0 {
		p.position.Y = 0
	}

	// player rotate
	if rl.IsKeyDown(rl.KeyQ) {
		p.angle -= float32(p.speed) * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyR) || rl.IsKeyDown(rl.KeyE) { // I included E as I don't like using R to rotate
		p.angle += float32(p.speed) * rl.GetFrameTime()
	}
}
