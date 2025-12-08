package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Blocker struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Color rl.Color
}

func NewBlocker(px, py, sx, sy float32, Color rl.Color) Blocker {
	return Blocker{Pos: rl.NewVector2(px, py), Size: rl.NewVector2(sx, sy), Color: Color}
}

func (bl Blocker) DrawBlocker() {
	rl.DrawRectangle(int32(bl.Pos.X), int32(bl.Pos.Y), int32(bl.Size.X), int32(bl.Size.Y), bl.Color)
}

func CheckCollision(physicsBody *PhysicsBody, blocker Blocker) {
	pbRec := physicsBody.GetPlayerRectangle()
	if rl.CheckCollisionRecs(pbRec, rl.NewRectangle(blocker.Pos.X, blocker.Pos.Y, blocker.Size.X, blocker.Size.Y)) {
		if pbRec.Y+physicsBody.Scale.Y > blocker.Pos.Y && pbRec.Y+physicsBody.Scale.Y < blocker.Pos.Y+blocker.Size.Y && physicsBody.Vel.Y > 0 && blocker.Size.Y < 300 { //creature over box
			//fmt.Println("Y Collision 1")
			physicsBody.Pos.Y = blocker.Pos.Y - physicsBody.Scale.Y/2
			physicsBody.Vel.Y = 0
			return
		}
		if pbRec.Y < blocker.Pos.Y+blocker.Size.Y && pbRec.Y > blocker.Pos.Y && physicsBody.Vel.Y < 0 && blocker.Size.Y < 300 { //creature under box
			//fmt.Println("Y Collision 2")
			physicsBody.Pos.Y = blocker.Pos.Y + blocker.Size.Y + physicsBody.Scale.Y/2
			physicsBody.Vel.Y = 0
			return
		}
		if pbRec.X < blocker.Pos.X+blocker.Size.X && pbRec.X+physicsBody.Scale.X > blocker.Pos.X+blocker.Size.X && physicsBody.Vel.X < 0 { // creature right of box
			//fmt.Println("X Collision 1")
			physicsBody.Pos.X = blocker.Pos.X + blocker.Size.X + physicsBody.Scale.X/4
			physicsBody.Vel.X = 0
			return
		}
		if pbRec.X+physicsBody.Scale.X > blocker.Pos.X && pbRec.X < blocker.Pos.X && physicsBody.Vel.X > 0 { // creature left of box
			//fmt.Println("X Collision 2")
			physicsBody.Pos.X = blocker.Pos.X - physicsBody.Scale.X/4
			physicsBody.Vel.X = 0
			return
		}
	}
}
