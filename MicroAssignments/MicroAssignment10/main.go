package main

import (
	"PhysicsEngine/ndphysics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Snowball")

	backgroundColor := rl.NewColor(47, 78, 128, 255)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	snowBall1 := ndphysics.NewProjectile(20, rl.NewVector2(200, 200), rl.NewVector2(200, 0))
	//snowBall2 := ndphysics.NewProjectile(20, rl.NewVector2(400, 200), rl.NewVector2(-20, 0))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)

		snowBall1.PhysicsUpdate()
		snowBall1.DrawProjectile()

		CheckEdges(&snowBall1)

		// snowBall2.PhysicsUpdate()
		// snowBall2.DrawProjectile()

		rl.EndDrawing()
	}
}

// probably didn't need to do all this, but I did so anyways
func CheckEdges(ball *ndphysics.Projectile) {
	var position rl.Vector2 = ball.GetPosition()
	var velocity rl.Vector2 = ball.GetVelocity()

	if (position.X+20 > 0 && position.X < 800) && (velocity.X != 0) {
		return
	} else if (position.Y+20 > 0 && position.Y < 450) && (velocity.Y != 0) {
		return
	}

	if velocity.X > 0 {
		ball.SetPosition(rl.NewVector2(-20, position.Y))
	} else if velocity.X < 0 {
		ball.SetPosition(rl.NewVector2(800, position.Y))
	} else if velocity.Y > 0 {
		ball.SetPosition(rl.NewVector2(position.X, -20))
	} else if velocity.Y < 0 {
		ball.SetPosition(rl.NewVector2(position.X, 450))
	}
}
