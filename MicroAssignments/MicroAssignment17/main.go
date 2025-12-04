package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "Gravity Box Example")
	defer rl.CloseWindow()

	box := Box{
		Pos:   rl.NewVector2(400, 350),
		Vel:   rl.NewVector2(0, 0),
		Size:  rl.NewVector2(50, 50),
		Color: rl.Red,
	}

	blockerDown := Blocker{
		Pos:   rl.NewVector2(300, 500), // Position of the blocker
		Size:  rl.NewVector2(200, 100),
		Color: rl.Gray,
	}
	blockerLeft := Blocker{
		Pos:   rl.NewVector2(200, 300),
		Size:  rl.NewVector2(100, 200),
		Color: rl.Gray,
	}
	blockerRight := Blocker{
		Pos:   rl.NewVector2(500, 300),
		Size:  rl.NewVector2(100, 200),
		Color: rl.Gray,
	}
	blockerUp := Blocker{
		Pos:   rl.NewVector2(300, 200),
		Size:  rl.NewVector2(200, 100),
		Color: rl.Gray,
	}

	gravity := rl.NewVector2(0, 980)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		if rl.IsKeyPressed(rl.KeyW) { // up
			gravity.X = 0
			gravity.Y = -980
		} else if rl.IsKeyPressed(rl.KeyA) { // left
			gravity.X = -980
			gravity.Y = 0
		} else if rl.IsKeyPressed(rl.KeyS) { // down
			gravity.X = 0
			gravity.Y = 980
		} else if rl.IsKeyPressed(rl.KeyD) { // right
			gravity.X = 980
			gravity.Y = 0
		}
		box.ApplyGravity(gravity)
		box.UpdateBox()

		CheckCollision(&box, blockerUp)
		CheckCollision(&box, blockerLeft)
		CheckCollision(&box, blockerDown)
		CheckCollision(&box, blockerRight)

		box.DrawBox()
		blockerUp.DrawBlocker()
		blockerLeft.DrawBlocker()
		blockerDown.DrawBlocker()
		blockerRight.DrawBlocker()

		rl.EndDrawing()
	}
}
