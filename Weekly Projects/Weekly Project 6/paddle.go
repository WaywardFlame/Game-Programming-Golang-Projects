package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Paddle struct {
	position  rl.Vector2
	velocity  rl.Vector2
	speed     float32
	direction int32
}

func (p *Paddle) VelocityTick() {
	if p.direction == 1 || p.direction == -1 {
		p.velocity.X = p.speed * float32(p.direction)
		adjustedVel := rl.Vector2Scale(p.velocity, rl.GetFrameTime())
		p.position = rl.Vector2Add(p.position, adjustedVel)
	} else if p.direction == 0 {
		p.velocity.X = 0
	}

	if p.position.X < 0 {
		p.position.X = 0
		p.direction = 0
	}
	if p.position.X > 700 {
		p.position.X = 700
		p.direction = 0
	}
}
