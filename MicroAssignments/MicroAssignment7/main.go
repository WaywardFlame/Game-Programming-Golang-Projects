package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	var playerX float32 = 25
	var playerY float32 = 100
	var playerSpeed float32 = 400
	var playerSize float32 = 50

	rl.InitWindow(800, 450, "Micro Assignment 7")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(int32(playerX), int32(playerY), int32(playerSize), int32(playerSize), rl.White)

		// keyboard input
		if rl.IsKeyDown(rl.KeyW) {
			playerY -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) {
			playerY += playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyA) {
			playerX -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyD) {
			playerX += playerSpeed * rl.GetFrameTime()
		}

		// check window edge
		if int32(playerX+playerSize) > 800 {
			playerX = 750
		}
		if int32(playerX) < 0 {
			playerX = 0
		}
		if int32(playerY+playerSize) > 450 {
			playerY = 400
		}
		if int32(playerY) < 0 {
			playerY = 0
		}

		rl.EndDrawing()
	}
}
