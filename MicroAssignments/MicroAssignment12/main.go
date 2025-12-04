package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Brian Statom - Micro Assignment 12")
	defer rl.CloseWindow()
	rl.InitAudioDevice()

	explosion := rl.LoadSound("audio/8 bit explosion.wav")
	carCrash := rl.LoadSound("audio/car crash.wav")
	mechaSound := rl.LoadSound("audio/mecha sound.wav")
	phoneBeep := rl.LoadSound("audio/phone beep.mp3")
	schoolBell := rl.LoadSound("audio/school bell.wav")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		rl.DrawText("Press 1, 2, 3, 4, or 5 to play a sound", 210, 200, 20, rl.White)
		if rl.IsKeyPressed(rl.KeyOne) {
			rl.PlaySound(explosion)
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			rl.PlaySound(carCrash)
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			rl.PlaySound(mechaSound)
		}
		if rl.IsKeyPressed(rl.KeyFour) {
			rl.PlaySound(phoneBeep)
		}
		if rl.IsKeyPressed(rl.KeyFive) {
			rl.PlaySound(schoolBell)
		}
		rl.EndDrawing()
	}
}
