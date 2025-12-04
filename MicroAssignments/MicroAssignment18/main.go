package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	idleSheet := rl.LoadTexture("sprites/idle.png")
	walkSheet := rl.LoadTexture("sprites/walk.png")
	rotateSheet := rl.LoadTexture("sprites/rotate.png")

	playerCreature := NewCreature(rl.NewVector2(400, 200), idleSheet, walkSheet, rotateSheet)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		HandleCreatureInput(&playerCreature)
		playerCreature.UpdateCreature()

		rl.EndDrawing()
	}
}

func HandleCreatureInput(creature *Creature) {
	direction := 0
	if rl.IsKeyDown(rl.KeyD) {
		direction = 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		direction = -1
	}
	if rl.IsKeyDown(rl.KeyR) {
		direction = 2
	}
	creature.MoveCreature(direction)
	if direction == 1 || direction == -1 {
		creature.AnimationStateMachine.ChangeState(WALKSTATE)
	} else if direction == 2 {
		creature.AnimationStateMachine.ChangeState(ROTATESTATE)
	} else {
		creature.AnimationStateMachine.ChangeState(IDLESTATE)
	}
}
