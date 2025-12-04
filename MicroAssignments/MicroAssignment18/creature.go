package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	*Transform
	PhysicsBody
	AnimationStateMachine StateMachine
	MoveSpeed             float32
}

func NewCreature(newPos rl.Vector2, idleSheet, walkSheet rl.Texture2D, rotateSheet rl.Texture2D) Creature {
	newTransform := NewTransform(newPos)

	newIdleAnimation := NewAnimation(newTransform, idleSheet, .2, IDLESTATE)
	newWalkAnimation := NewAnimation(newTransform, walkSheet, .1, WALKSTATE)
	newRotateAnimation := NewAnimation(newTransform, rotateSheet, .1, ROTATESTATE)

	newCreature := Creature{
		Transform:             newTransform,
		PhysicsBody:           PhysicsBody{Transform: newTransform},
		AnimationStateMachine: NewStateMachine(&newIdleAnimation),
		MoveSpeed:             100,
	}

	newCreature.Scale = 8

	newCreature.AnimationStateMachine.AddState(&newWalkAnimation)
	newCreature.AnimationStateMachine.AddState(&newRotateAnimation)
	return newCreature
}

func (c *Creature) UpdateCreature() {
	c.ApplyVel()
	c.AnimationStateMachine.Tick()
}

func (c *Creature) MoveCreature(direction int) {
	if direction == 1 {
		c.Flip = 1
	}
	if direction == -1 {
		c.Flip = -1
	}
	if direction != 2 {
		c.Vel.X = c.MoveSpeed * float32(direction)
	}
}
