package main

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawTextureEz(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)
	rl.DrawTexturePro(texture, sourceRect, destRect, origin, angle, color)
}

func main() {
	rl.InitWindow(800, 450, "Micro Assignment 8")
	defer rl.CloseWindow()

	sprite := rl.LoadTexture("textures/sprite.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		DrawTextureEz(sprite, rl.NewVector2(200, 225), -45, 2, rl.NewColor(225, 15, 15, 255))
		DrawTextureEz(sprite, rl.NewVector2(375, 225), 180, 4, rl.NewColor(15, 225, 15, 255))
		DrawTextureEz(sprite, rl.NewVector2(625, 225), 45, 6, rl.NewColor(15, 15, 225, 255))
		rl.EndDrawing()
	}
}
