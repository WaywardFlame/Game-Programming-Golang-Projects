package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Tree struct {
	pos rl.Vector2
}

func NewTree(newPos rl.Vector2) Tree {
	newTree := Tree{pos: newPos}
	return newTree
}

func (t Tree) DrawTree() {
	rl.DrawCircle(int32(t.pos.X), int32(t.pos.Y), 30, rl.Lime)
}
