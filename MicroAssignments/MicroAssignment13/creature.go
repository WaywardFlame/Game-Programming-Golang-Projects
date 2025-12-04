package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	pos   rl.Vector2
	speed float32
}

func NewCreature(newPos rl.Vector2) Creature {
	newCreature := Creature{pos: newPos, speed: 100}
	return newCreature
}

func (c Creature) DrawCreature() {
	rl.DrawCircle(int32(c.pos.X), int32(c.pos.Y), 20, rl.White)
}

func (c *Creature) MoveCreature(dir rl.Vector2) {

	c.pos = rl.Vector2Add(c.pos, rl.Vector2Scale(dir, c.speed*rl.GetFrameTime()))
}

func (c *Creature) MoveCreatureWithCamera(input rl.Vector2, angle float32) {
	rad := float64((angle) * (math.Pi / 180))

	upVec := rl.NewVector2(float32(math.Sin(rad)), float32(math.Cos(rad)))

	c.pos = rl.Vector2Add(c.pos, rl.Vector2Scale(upVec, input.Y*c.speed*rl.GetFrameTime()))

	// radRight := float64(angle+90) * (math.Pi / 180)
	// rightVec := rl.NewVector2(float32(math.Sin(radRight)), float32(math.Cos(radRight)))
	rightVec := rl.NewVector2(upVec.Y, -upVec.X)
	c.pos = rl.Vector2Add(c.pos, rl.Vector2Scale(rightVec, input.X*c.speed*rl.GetFrameTime()))
}
