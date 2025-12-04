package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	IDLESTATE  = "idle"
	WALKSTATE  = "walk"
	JUMPSTATE  = "jump"
	PUNCHSTATE = "punch"
	KICKSTATE  = "kick"
	BLOCKSTATE = "block"
)

type Animation struct {
	*Transform
	*Actions
	Color        rl.Color
	SpriteSheet  rl.Texture2D
	MaxIndex     int
	CurrentIndex int
	Timer        float32
	SwitchTime   float32
	Name         string
}

func (a *Animation) TickState() {
	a.UpdateTime()
	a.DrawAnimation()
}

func (a *Animation) GetName() string {
	return a.Name
}

func (a *Animation) ResetTime() {
	a.Timer = 0
}

func NewAnimation(newTransform *Transform, newActions *Actions, color rl.Color, newSheet rl.Texture2D, newTime float32, newName string) Animation {
	fmt.Println(newSheet.Width, newSheet.Height)
	frames := int(newSheet.Width / 128)
	newAnimation := Animation{
		Transform:    newTransform,
		Actions:      newActions,
		Color:        color,
		SpriteSheet:  newSheet,
		MaxIndex:     frames - 1,
		CurrentIndex: 0,
		Timer:        0,
		SwitchTime:   newTime,
		Name:         newName,
	}
	return newAnimation
}

func (a *Animation) UpdateTime() {
	a.Timer += rl.GetFrameTime()
	if a.Timer > a.SwitchTime {
		a.Timer = 0
		a.CurrentIndex++
	}

	if a.CurrentIndex > a.MaxIndex {
		a.CurrentIndex = 0
		if a.isPunching {
			a.isPunching = false
		}
		if a.isKicking {
			a.isKicking = false
		}
		if a.isJumping {
			a.isJumping = false
		}
		if a.isBlocking {
			a.isBlocking = false
		}
	}
}

func (a Animation) DrawAnimation() {
	sourceRect := rl.NewRectangle(float32(128*a.CurrentIndex), 0, 128*float32(a.Flip), 64)
	destRect := rl.NewRectangle(a.Pos.X, a.Pos.Y, float32(a.Scale.X), float32(a.Scale.Y))
	origin := rl.NewVector2(a.Scale.X/2, a.Scale.Y/2)
	rl.DrawTexturePro(a.SpriteSheet, sourceRect, destRect, origin, 0, a.Color)
}
