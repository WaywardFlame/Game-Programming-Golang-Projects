package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	mouseOver bool
	position  rl.Vector2
	texture   rl.Texture2D
	checked   bool
}

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 14")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	button := Button{mouseOver: false, position: rl.NewVector2(380, 205), texture: rl.LoadTexture("textures/EmptyBox.png"), checked: false}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)
		rl.DrawTexture(button.texture, int32(button.position.X), int32(button.position.Y), rl.White)

		// check mouse position
		mousePos := rl.GetMousePosition()
		if mousePos.X > button.position.X && mousePos.X < button.position.X+20 {
			if mousePos.Y > button.position.Y && mousePos.Y < button.position.Y+20 {
				button.mouseOver = true
			} else {
				button.mouseOver = false
			}
		} else {
			button.mouseOver = false
		}

		// check mouse click
		if button.mouseOver && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
			if button.checked {
				button.texture = rl.LoadTexture("textures/EmptyBox.png")
				button.checked = false
			} else {
				button.texture = rl.LoadTexture("textures/CheckBox.png")
				button.checked = true
			}
		}

		rl.EndDrawing()
	}
}
