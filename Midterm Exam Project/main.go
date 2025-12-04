package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerCircle struct {
	pos    rl.Vector2
	radius int
	color  rl.Color
}

type Mine struct {
	pos    rl.Vector2
	radius int
	armed  bool
}

func (m *Mine) checkPlayerPosition(pc PlayerCircle) {
	if pc.pos.X+float32(pc.radius) < m.pos.X-float32(m.radius) || pc.pos.X-float32(pc.radius) > m.pos.X+float32(m.radius) {
		m.armed = true
	}
	if pc.pos.Y+float32(pc.radius) < m.pos.Y-float32(m.radius) || pc.pos.Y-float32(pc.radius) > m.pos.Y+float32(m.radius) {
		m.armed = true
	}
}

func (m *Mine) checkGameOver(pc PlayerCircle) bool {
	if rl.Vector2Distance(m.pos, pc.pos) <= float32(m.radius+pc.radius) {
		return true
	}
	return false
}

func main() {
	rl.InitWindow(1280, 720, "Brian Statom - Midterm Exam Project")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	pc := PlayerCircle{radius: 128, color: rl.Blue}
	mines := make([]Mine, 0, 20)
	var score int = 0
	var gameOver bool = false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if gameOver {
			rl.ClearBackground(rl.Red)
			rl.DrawText("Press R to restart", 0, 0, 40, rl.Black)
			rl.DrawText("Score: "+strconv.Itoa(score), 0, 40, 30, rl.Black)
			if rl.IsKeyPressed(rl.KeyR) {
				score = 0
				gameOver = false
				mines = make([]Mine, 0, 20)
			}
		} else {
			rl.ClearBackground(rl.Black)
			pc.pos = rl.GetMousePosition()
			rl.DrawCircle(int32(pc.pos.X), int32(pc.pos.Y), float32(pc.radius), pc.color)
			if rl.IsKeyPressed(rl.KeySpace) {
				mines = append(mines, Mine{pos: pc.pos, radius: 64, armed: false})
				score++
			}
			rl.DrawText("Score: "+strconv.Itoa(score), 0, 0, 40, rl.White)
			for i := 0; i < len(mines); i++ {
				if !mines[i].armed {
					mines[i].checkPlayerPosition(pc)
				}
				if mines[i].armed {
					rl.DrawCircle(int32(mines[i].pos.X), int32(mines[i].pos.Y), float32(mines[i].radius), rl.Red)
					gameOver = mines[i].checkGameOver(pc)
					if gameOver {
						break
					}
				} else {
					rl.DrawCircle(int32(mines[i].pos.X), int32(mines[i].pos.Y), float32(mines[i].radius), rl.Gray)
				}
			}
		}

		rl.EndDrawing()
	}
}
