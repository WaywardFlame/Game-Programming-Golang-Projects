package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// functions that are specific to player projectiles

func newProjectile(position rl.Vector2, angle float64, speed float32, texture rl.Texture2D) PhysicsBody {
	rad := float64((angle) * (math.Pi / 180))
	// direction * speed = velocity
	direction := rl.NewVector2(float32(math.Sin(rad)), -float32(math.Cos(rad))) // get direction
	vel := rl.Vector2Scale(direction, float32(speed+100))
	projectile := newPhysicsBody(position, vel, 8, texture, true)
	projectile.angle = float32(angle)
	return projectile
}

func (pb *PhysicsBody) ProjectileUpdate() {
	pb.position = rl.Vector2Add(pb.position, rl.Vector2Scale(pb.velocity, rl.GetFrameTime()))
}
