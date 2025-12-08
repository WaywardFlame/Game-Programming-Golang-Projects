package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(600, 900, "Weekly Project 8 - Laser Reflection Game")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	var FPS int = 60
	rl.SetTargetFPS(int32(FPS))

	textures := []rl.Texture2D{
		rl.LoadTexture("textures/LaserBall.png"),
		rl.LoadTexture("textures/LaserRect.png"),
		rl.LoadTexture("textures/playerTurret.png"),
	}
	sounds := []rl.Sound{
		rl.LoadSound("audio/333785__aceofspadesproduc100__8-bit-failure-sound.wav"), // failure sound
		rl.LoadSound("audio/456590__bumpelsnake__laser1a.wav"),                      // laser sound
		rl.LoadSound("audio/531512__cogfirestudios__positive-blip-effect.wav"),      // success sound
	}
	gameData := GameData{Level: 0, Player: newTurret(rl.NewVector2(300, 780), 50, 0),
		Textures: textures, Sounds: sounds, Initialized: false}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawRectangle(20, 20, 560, 860, rl.Black)

		if gameData.Level == 0 {
			rl.DrawText("Welcome to the Laser Reflection Game", 100, 40, 20, rl.White)
			rl.DrawText("Aim your turret to carefully reflect", 100, 80, 20, rl.White)
			rl.DrawText("your laser off mirrors (cyan blocks).", 100, 100, 20, rl.White)
			rl.DrawText("Rotate your turret with Q and E.", 100, 120, 20, rl.White)
			rl.DrawText("Press W to fire a laser from turret.", 100, 140, 20, rl.White)
			rl.DrawText("Hit the green block to succeed.", 100, 160, 20, rl.White)
			rl.DrawText("Hit the red blocks to open paths.", 100, 180, 20, rl.White)
			rl.DrawText("You have 5 shots per level, and restart", 100, 220, 20, rl.White)
			rl.DrawText("From the first level of the game.", 100, 240, 20, rl.White)
			rl.DrawText("Press Space to begin the game.", 100, 280, 20, rl.White)
			rl.DrawText("You can practice your shots here.", 100, 300, 20, rl.White)

			gameData.drawMainMenu()

			if rl.IsKeyPressed(rl.KeySpace) {
				rl.PlaySound(gameData.Sounds[2])
				gameData.Initialized = false
				gameData.GameOvers = 0
				gameData.LasersFired = 0
				gameData.initializeLevelOne()
			}
		} else if gameData.Level == 1 {
			gameData.initializeLevelOne()
		} else if gameData.Level == 2 {
			gameData.initializeLevelTwo()
		} else if gameData.Level == 3 {
			gameData.initializeLevelThree()
		} else if gameData.Level == 4 { // end screen
			rl.DrawText("Congrats! You've reached the end", 100, 40, 20, rl.White)
			rl.DrawText("of the game. Good job!", 100, 60, 20, rl.White)
			rl.DrawText("Press Space to return to the main menu.", 100, 100, 20, rl.White)
			rl.DrawText("Lasers Fired: "+strconv.Itoa(gameData.LasersFired), 100, 140, 20, rl.White)
			rl.DrawText("Game Overs: "+strconv.Itoa(gameData.GameOvers), 100, 160, 20, rl.White)
			if rl.IsKeyPressed(rl.KeySpace) {
				rl.PlaySound(gameData.Sounds[2])
				gameData.Initialized = false
				gameData.Level = 0
			}
		}

		rl.EndDrawing()
	}
}
