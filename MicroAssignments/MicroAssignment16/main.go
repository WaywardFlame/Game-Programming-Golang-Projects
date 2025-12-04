package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Alive    bool
	Level    int
	Size     float32
	Position rl.Vector2
	Color    rl.Color
}

type GameData struct {
	Player  Creature
	Enemies []Creature
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

	player := Creature{Alive: true, Level: 1, Size: 50, Position: generatePlayerPosition(), Color: playerColor}
	var enemies []Creature
	for i := 0; i < 5; i++ {
		enemies = append(enemies, Creature{Alive: true, Level: i + 1, Size: 50, Color: enemyColor})
		enemies[i].Position = generateEnemyPosition(player.Position, enemies, i)
	}

	var game GameData

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if gameWin {
			rl.ClearBackground(rl.Blue)
			rl.DrawText("Game Over: You WIN", 150, 185, 50, rl.White)
			if rl.IsKeyPressed(rl.KeyR) {
				player = Creature{Alive: true, Level: 1, Size: 50, Position: generatePlayerPosition(), Color: playerColor}
				for i := 0; i < 5; i++ {
					enemies[i] = Creature{Alive: true, Level: i + 1, Size: 50, Color: enemyColor}
					enemies[i].Position = generateEnemyPosition(player.Position, enemies, i)
				}
				gameWin = false
			}
		} else if gameLose {
			rl.ClearBackground(rl.Red)
			rl.DrawText("Game Over: You LOSE", 145, 185, 50, rl.White)
			if rl.IsKeyPressed(rl.KeyR) {
				player = Creature{Alive: true, Level: 1, Size: 50, Position: generatePlayerPosition(), Color: playerColor}
				for i := 0; i < 5; i++ {
					enemies[i] = Creature{Alive: true, Level: i + 1, Size: 50, Color: enemyColor}
					enemies[i].Position = generateEnemyPosition(player.Position, enemies, i)
				}
				gameLose = false
			}
		} else {
			rl.ClearBackground(rl.LightGray)
			if rl.IsKeyPressed(rl.KeyK) {
				game = GameData{Player: player, Enemies: enemies}
				game.Save("savedata.json")
			}
			if rl.IsKeyPressed(rl.KeyL) {
				game.Load("savedata.json")
				player = game.Player
				enemies = game.Enemies
			}

			rl.DrawTextureEx(sprite, player.Position, 0, 1.5625, playerColor) // draw player
			rl.DrawText(strconv.Itoa(player.Level), int32(player.Position.X), int32(player.Position.Y), 20, rl.White)
			for i := 0; i < 5; i++ {
				rl.DrawTextureEx(sprite, enemies[i].Position, 0, 1.5625, enemies[i].Color) // draw enemies
				if enemies[i].Alive {
					rl.DrawText(strconv.Itoa(enemies[i].Level), int32(enemies[i].Position.X), int32(enemies[i].Position.Y), 20, rl.White)
				}
			}

			// player movement, prevent exiting screen
			if rl.IsKeyPressed(rl.KeyW) {
				player.Position.Y -= player.Size
				if player.Position.Y < 0 {
					player.Position.Y = 0
				}
			} else if rl.IsKeyPressed(rl.KeyA) {
				player.Position.X -= player.Size
				if player.Position.X < 0 {
					player.Position.X = 0
				}
			} else if rl.IsKeyPressed(rl.KeyS) {
				player.Position.Y += player.Size
				if player.Position.Y > 400 {
					player.Position.Y = 400
				}
			} else if rl.IsKeyPressed(rl.KeyD) {
				player.Position.X += player.Size
				if player.Position.X > 750 {
					player.Position.X = 750
				}
			}

			// check for overlap with enemies
			var deadEnemies int = 0
			for i := 0; i < 5; i++ {
				if !enemies[i].Alive {
					deadEnemies++
					continue
				}
				// to hopefully account for floating point weirdness
				tempX := player.Position.X - enemies[i].Position.X
				tempY := player.Position.Y - enemies[i].Position.Y
				if (tempX >= -1 && tempX <= 1) && (tempY >= -1 && tempY <= 1) {
					if player.Level >= enemies[i].Level {
						player.Level += enemies[i].Level
						enemies[i].Alive = false
						enemies[i].Color = deadColor
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
		// check against player Position
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
		for k := 0; !notOnPOS; k++ { // check against enemy Positions
			// to hopefully account for floating weirdness
			tempX := enemies[k].Position.X - tempVector.X
			tempY := enemies[k].Position.Y - tempVector.Y
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

func (g *GameData) Save(filename string) error {
	data, err := json.MarshalIndent(g, "", "	")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (g *GameData) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, g)
}
