package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1280, 720, "Micro Assignment 6")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawRectangle(0, 0, 60, 60, rl.Red)
		rl.DrawRectangle(1220, 0, 60, 60, rl.Green)
		rl.DrawRectangle(1220, 660, 60, 60, rl.Blue)
		rl.DrawRectangle(0, 660, 60, 60, rl.Yellow)
		rl.DrawRectangle(610, 330, 60, 60, rl.Black)
		rl.EndDrawing()
	}
}
