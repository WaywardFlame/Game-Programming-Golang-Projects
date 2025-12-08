package main

import (
	"math/rand/v2"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var gameOver bool = false

	// player variables
	var playerX float32 = 100
	var playerY float32 = 100
	var playerSpeed float32 = 300
	var playerSize float32 = 50
	var playerScore int = 0
	var scoreAwarded bool = false

	// pipe variables
	var pipeX float32 = -100
	var pipeY float32
	var pipeSpeed float32 = 300
	var pipeSize float32 = 50

	rl.InitWindow(800, 450, "Weekly Project 4 - Flappy Bird Clone")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if !gameOver {
			rl.ClearBackground(rl.Black)
			rl.DrawRectangle(int32(playerX), int32(playerY), int32(playerSize), int32(playerSize), rl.Blue)

			// keyboard input
			if rl.IsKeyDown(rl.KeyW) {
				playerY -= playerSpeed * rl.GetFrameTime()
			}
			if rl.IsKeyDown(rl.KeyS) {
				playerY += playerSpeed * rl.GetFrameTime()
			}

			// check window edge
			if int32(playerY+playerSize) > 450 {
				playerY = 400
			}
			if int32(playerY) < 0 {
				playerY = 0
			}

			// draw pipes
			// check if off left side of screen first
			if pipeX+pipeSize < 0 {
				pipeX = 850
				pipeY = float32(rand.IntN(276) + 50)
				scoreAwarded = false
			}
			// now draw the pipes moving
			rl.DrawRectangle(int32(pipeX), 0, int32(pipeSize), int32(pipeY), rl.Green)
			rl.DrawRectangle(int32(pipeX), int32(pipeY+100), int32(pipeSize), 500, rl.Green)
			pipeX -= pipeSpeed * rl.GetFrameTime()

			// check player intersecting with pipe
			if int32(pipeX) <= int32(playerX+playerSize) && !(int32(pipeX+pipeSize) < int32(playerX)) {
				if int32(playerY) < int32(pipeY) {
					gameOver = true
				} else if int32(playerY) > int32(pipeY+100) {
					gameOver = true
				} else if int32(playerY+playerSize) > int32(pipeY+100) {
					gameOver = true
				}
			}

			if !scoreAwarded && playerX > pipeX+pipeSize {
				playerScore++
				scoreAwarded = true
			}
			rl.DrawText("Score: "+strconv.Itoa(playerScore), 0, 0, 20, rl.White)

		} else {
			rl.ClearBackground(rl.Red)
			rl.DrawText("Game Over - Press 'R' to restart", 240, 215, 20, rl.White)

			if rl.IsKeyDown(rl.KeyR) {
				gameOver = false
				pipeX = -100
				playerX = 100
				playerY = 100
				playerScore = 0
			}
		}
		rl.EndDrawing()
	}
}
