package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Blocker struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Color rl.Color
}

func (bl Blocker) DrawBlocker() {
	rl.DrawRectangle(int32(bl.Pos.X), int32(bl.Pos.Y), int32(bl.Size.X), int32(bl.Size.Y), bl.Color)
}

func CheckCollision(box *Box, blocker Blocker) {
	if rl.CheckCollisionRecs( //Raylib let's us quickly check overlap with the rectangle class.
		rl.NewRectangle(box.Pos.X, box.Pos.Y, box.Size.X, box.Size.Y),
		rl.NewRectangle(blocker.Pos.X, blocker.Pos.Y, blocker.Size.X, blocker.Size.Y),
	) {
		if box.Pos.Y+box.Size.Y > blocker.Pos.Y && box.Vel.Y > 0 { //now check which side to stop the velocity
			box.Pos.Y = blocker.Pos.Y - box.Size.Y //move box in case of overlap
			box.Vel.Y = 0                          //stop the box from moving further
		}
		if box.Pos.Y < blocker.Pos.Y+blocker.Size.Y && box.Vel.Y < 0 {
			box.Pos.Y = blocker.Pos.Y + blocker.Size.Y
			box.Vel.Y = 0
		}
		if box.Pos.X+box.Size.X > blocker.Pos.X && box.Vel.X > 0 {
			box.Pos.X = blocker.Pos.X - box.Size.X
			box.Vel.X = 0
		}
		if box.Pos.X < blocker.Pos.X+blocker.Size.X && box.Vel.X < 0 {
			box.Pos.X = blocker.Pos.X + blocker.Size.X
			box.Vel.X = 0
		}
	}
}
