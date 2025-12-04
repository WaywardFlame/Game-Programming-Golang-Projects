package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Brian Statom - Project 6 Breakout")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	paddle := Paddle{position: rl.NewVector2(350, 375), velocity: rl.NewVector2(0, 0), speed: 300, direction: 0}
	ball := Ball{position: rl.NewVector2(400, 365), velocity: rl.NewVector2(0, 0), radius: 10, launched: false}
	blocks := createBlocks() // remeber 62 by 62 size
	var gameOver bool = false
	var blocksHit int = 0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkBlue)

		if gameOver {
			paddle = Paddle{position: rl.NewVector2(350, 375), velocity: rl.NewVector2(0, 0), speed: 300}
			ball = Ball{position: rl.NewVector2(400, 365), velocity: rl.NewVector2(0, 0), radius: 10, launched: false}
			blocks = createBlocks() // inefficient, but whatever
			paddle.direction = 0
			gameOver = false
			blocksHit = 0
		} else {
			// draw paddle
			rl.DrawRectangle(int32(paddle.position.X), int32(paddle.position.Y), 100, 5, rl.LightGray)
			if rl.IsKeyDown(rl.KeyA) {
				paddle.direction = -1
				paddle.VelocityTick()
			} else if rl.IsKeyDown(rl.KeyD) {
				paddle.direction = 1
				paddle.VelocityTick()
			} else {
				paddle.direction = 0
				paddle.VelocityTick()
			}

			// draw blocks
			for i := 0; i < 3; i++ {
				for k := 0; k < 10; k++ {
					if !blocks[i][k].visible {
						continue
					}
					rl.DrawRectangle(int32(blocks[i][k].position.X), int32(blocks[i][k].position.Y), 62, 62, rl.Orange)
				}
			}

			// draw ball
			if rl.IsKeyPressed(rl.KeySpace) && !ball.launched {
				ball.launched = true
				if paddle.direction == -1 {
					ball.velocity = rl.NewVector2(-300, -300)
				} else if paddle.direction == 1 {
					ball.velocity = rl.NewVector2(300, -300)
				} else if paddle.direction == 0 {
					ball.velocity = rl.NewVector2(0, -300)
				}
			}
			if ball.launched {
				ball.VelocityTick()
				ball.CheckForWalls()
				ball.CheckForPaddle(paddle, blocksHit)
				blocksHit = ball.CheckForBlock(&blocks, blocksHit)
				if blocksHit == 30 {
					gameOver = true
				}
				rl.DrawCircle(int32(ball.position.X), int32(ball.position.Y), ball.radius, rl.White)
				if ball.position.Y >= 450 {
					gameOver = true
				}
			} else {
				if paddle.direction == -1 {
					ball.position.X -= paddle.speed * rl.GetFrameTime()
				} else if paddle.direction == 1 {
					ball.position.X += paddle.speed * rl.GetFrameTime()
				}

				if ball.position.X < paddle.position.X+40 || ball.position.X > paddle.position.X+60 {
					ball.position.X = paddle.position.X + 50
				}
				rl.DrawCircle(int32(ball.position.X), int32(ball.position.Y), ball.radius, rl.White)
			}
		}
		rl.EndDrawing()
	}
}
