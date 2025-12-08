package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	*Transform
	PhysicsBody
	*Actions
	AnimationStateMachine StateMachine
	MoveSpeed             float32
	Health                int
	HitTimer              float32
	Victories             int
}

type Actions struct {
	isPunching bool
	isKicking  bool
	isJumping  bool
	isBlocking bool
}

func NewCreature(newPos rl.Vector2, SpriteSheets []rl.Texture2D, playerColor rl.Color) Creature {
	newTransform := NewTransform(newPos)
	var newActions *Actions = &Actions{isPunching: false, isKicking: false, isJumping: false, isBlocking: false}

	newIdleAnimation := NewAnimation(newTransform, newActions, playerColor, SpriteSheets[0], .2, IDLESTATE)
	newWalkAnimation := NewAnimation(newTransform, newActions, playerColor, SpriteSheets[1], .1, WALKSTATE)
	newJumpAnimation := NewAnimation(newTransform, newActions, playerColor, SpriteSheets[2], .2, JUMPSTATE)
	newPunchAnimation := NewAnimation(newTransform, newActions, playerColor, SpriteSheets[3], .1, PUNCHSTATE)
	newKickAnimation := NewAnimation(newTransform, newActions, playerColor, SpriteSheets[4], .1, KICKSTATE)
	newBlockAnimation := NewAnimation(newTransform, newActions, playerColor, SpriteSheets[5], .1, BLOCKSTATE)

	newCreature := Creature{
		Transform:             newTransform,
		PhysicsBody:           PhysicsBody{Transform: newTransform},
		Actions:               newActions,
		AnimationStateMachine: NewStateMachine(&newIdleAnimation),
		MoveSpeed:             300,
		Health:                100,
		HitTimer:              0,
		Victories:             0,
	}

	newCreature.Scale = rl.NewVector2(256, 128)

	newCreature.AnimationStateMachine.AddState(&newWalkAnimation)
	newCreature.AnimationStateMachine.AddState(&newJumpAnimation)
	newCreature.AnimationStateMachine.AddState(&newPunchAnimation)
	newCreature.AnimationStateMachine.AddState(&newKickAnimation)
	newCreature.AnimationStateMachine.AddState(&newBlockAnimation)

	return newCreature
}

func (c *Creature) UpdateCreature() {
	c.ApplyVel()
	c.AnimationStateMachine.Tick()
}

func (c *Creature) Jump() {
	c.Vel.Y = -500
}

func (c *Creature) CanJump(blockers []Blocker) bool {
	jumpCheckRect1 := rl.NewRectangle(5+c.Pos.X-c.Scale.X/4, c.Pos.Y+c.Scale.Y/2, 10, 10)
	jumpCheckRect2 := rl.NewRectangle(113+c.Pos.X-c.Scale.X/4, c.Pos.Y+c.Scale.Y/2, 10, 10)
	// rl.DrawRectangleRec(jumpCheckRect1, rl.Green)
	// rl.DrawRectangleRec(jumpCheckRect2, rl.Green)
	for _, b := range blockers {
		block := rl.NewRectangle(b.Pos.X, b.Pos.Y, b.Size.X, b.Size.Y)
		if rl.CheckCollisionRecs(jumpCheckRect1, block) {
			return true
		} else if rl.CheckCollisionRecs(jumpCheckRect2, block) {
			return true
		}
	}
	return false
}

func (c *Creature) MoveCreature(direction int) {
	if direction == 1 {
		c.Flip = 1
	}
	if direction == -1 {
		c.Flip = -1
	}
	c.Vel.X = float32(direction * int(c.MoveSpeed))
}
