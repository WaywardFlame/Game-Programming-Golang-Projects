package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	Position   rl.Vector2
	Angle      float64
	Velocity   rl.Vector2
	TextBall   rl.Texture2D
	TextRect   rl.Texture2D
	TrailTimer int
}

func newProjectile(position rl.Vector2, angle float64, speed float32, tex1 rl.Texture2D, tex2 rl.Texture2D) Projectile {
	rad := float64((angle) * (math.Pi / 180))
	// direction * speed = velocity
	direction := rl.NewVector2(float32(math.Sin(rad)), -float32(math.Cos(rad))) // get direction
	vel := rl.Vector2Scale(direction, float32(speed+100))
	projectile := Projectile{Position: position, Angle: angle, Velocity: vel, TextBall: tex1, TextRect: tex2, TrailTimer: 0}
	return projectile
}

func (pb *Projectile) ProjectileUpdate() {
	pb.Position = rl.Vector2Add(pb.Position, rl.Vector2Scale(pb.Velocity, rl.GetFrameTime()))
}
