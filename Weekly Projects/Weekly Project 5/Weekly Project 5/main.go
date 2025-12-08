package main

import (
	"math/rand/v2"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	alive    bool
	level    int
	size     float32
	position rl.Vector2
	color    rl.Color
}

func main() {
	rl.InitWindow(800, 450, "Weekly Project 5 - Turn Based Hero")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	var gameWin bool = false
	var gameLose bool = false
	sprite := rl.LoadTexture("textures/sprite.png")
	playerColor := rl.NewColor(0, 0, 255, 255)
	enemyColor := rl.NewColor(255, 0, 0, 255)
	deadColor := rl.NewColor(0, 0, 0, 0)

	player := Creature{alive: true, level: 1, size: 50, position: generatePlayerPosition(), color: playerColor}
	var enemies []Creature
	for i := 0; i < 5; i++ {
		enemies = append(enemies, Creature{alive: true, level: i + 1, size: 50, color: enemyColor})
		enemies[i].position = generateEnemyPosition(player.position, enemies, i)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if gameWin {
			rl.ClearBackground(rl.Blue)
			rl.DrawText("Game Over: You WIN", 150, 185, 50, rl.White)
			if rl.IsKeyPressed(rl.KeyR) {
				player = Creature{alive: true, level: 1, size: 50, position: generatePlayerPosition(), color: playerColor}
				for i := 0; i < 5; i++ {
					enemies[i] = Creature{alive: true, level: i + 1, size: 50, color: enemyColor}
					enemies[i].position = generateEnemyPosition(player.position, enemies, i)
				}
				gameWin = false
			}
		} else if gameLose {
			rl.ClearBackground(rl.Red)
			rl.DrawText("Game Over: You LOSE", 145, 185, 50, rl.White)
			if rl.IsKeyPressed(rl.KeyR) {
				player = Creature{alive: true, level: 1, size: 50, position: generatePlayerPosition(), color: playerColor}
				for i := 0; i < 5; i++ {
					enemies[i] = Creature{alive: true, level: i + 1, size: 50, color: enemyColor}
					enemies[i].position = generateEnemyPosition(player.position, enemies, i)
				}
				gameLose = false
			}
		} else {
			rl.ClearBackground(rl.LightGray)
			rl.DrawTextureEx(sprite, player.position, 0, 1.5625, playerColor) // draw player
			rl.DrawText(strconv.Itoa(player.level), int32(player.position.X), int32(player.position.Y), 20, rl.White)
			for i := 0; i < 5; i++ {
				rl.DrawTextureEx(sprite, enemies[i].position, 0, 1.5625, enemies[i].color) // draw enemies
				if enemies[i].alive {
					rl.DrawText(strconv.Itoa(enemies[i].level), int32(enemies[i].position.X), int32(enemies[i].position.Y), 20, rl.White)
				}
			}

			// player movement, prevent exiting screen
			if rl.IsKeyPressed(rl.KeyW) {
				player.position.Y -= player.size
				if player.position.Y < 0 {
					player.position.Y = 0
				}
			} else if rl.IsKeyPressed(rl.KeyA) {
				player.position.X -= player.size
				if player.position.X < 0 {
					player.position.X = 0
				}
			} else if rl.IsKeyPressed(rl.KeyS) {
				player.position.Y += player.size
				if player.position.Y > 400 {
					player.position.Y = 400
				}
			} else if rl.IsKeyPressed(rl.KeyD) {
				player.position.X += player.size
				if player.position.X > 750 {
					player.position.X = 750
				}
			}

			// check for overlap with enemies
			var deadEnemies int = 0
			for i := 0; i < 5; i++ {
				if !enemies[i].alive {
					deadEnemies++
					continue
				}
				// to hopefully account for floating point weirdness
				tempX := player.position.X - enemies[i].position.X
				tempY := player.position.Y - enemies[i].position.Y
				if (tempX >= -1 && tempX <= 1) && (tempY >= -1 && tempY <= 1) {
					if player.level >= enemies[i].level {
						player.level += enemies[i].level
						enemies[i].alive = false
						enemies[i].color = deadColor
					} else {
						gameLose = true
						break
					}
				}
			}
			if deadEnemies == 5 {
				gameWin = true
			}
		}
		rl.EndDrawing()
	}
}

func generatePlayerPosition() rl.Vector2 {
	x := rand.IntN(16) // 800 divided by 50
	y := rand.IntN(9)  // 450 divided by 50
	return rl.NewVector2(float32(x*50), float32(y*50))
}

func generateEnemyPosition(playerPOS rl.Vector2, enemies []Creature, i int) rl.Vector2 {
	var notOnPOS bool = false
	var tempVector rl.Vector2
	for !notOnPOS {
		tempVector = rl.NewVector2(float32(rand.IntN(16)*50), float32(rand.IntN(9)*50))
		// check against player position
		// to hopefully account for floating point weirdness
		tempX := playerPOS.X - tempVector.X
		tempY := playerPOS.Y - tempVector.Y
		if (tempX >= -1 && tempX <= 1) && (tempY >= -1 && tempY <= 1) {
			continue
		}
		if i == 0 { // if first enemy
			notOnPOS = true
			break
		}
		for k := 0; !notOnPOS; k++ { // check against enemy positions
			// to hopefully account for floating weirdness
			tempX := enemies[k].position.X - tempVector.X
			tempY := enemies[k].position.Y - tempVector.Y
			if (tempX >= -1 && tempX <= 1) && (tempY >= -1 && tempY <= 1) {
				break
			}
			if k == i {
				notOnPOS = true
			}
		}
	}
	return tempVector
}
