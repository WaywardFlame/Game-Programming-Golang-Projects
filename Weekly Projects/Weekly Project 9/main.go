package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// a lot of code was taken from AnimationFSM, created by
// the professor, Samuel Ang
// also from Platformer

func main() {
	rl.InitWindow(1280, 720, "Brian Statom - Project 9 - Fighting Game")
	rl.InitAudioDevice()
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	// load in sprite sheets
	SpriteSheets := []rl.Texture2D{
		rl.LoadTexture("textures/fighter_idle.png"),
		rl.LoadTexture("textures/fighter_walk.png"),
		rl.LoadTexture("textures/fighter_jump.png"),
		rl.LoadTexture("textures/fighter_punch.png"),
		rl.LoadTexture("textures/fighter_kick.png"),
		rl.LoadTexture("textures/fighter_block.png"),
	}

	// load in sound and music
	SoundEffects := []rl.Sound{
		rl.LoadSound("audio/176833__tapepusher__kick-2.wav"),       // kick sound
		rl.LoadSound("audio/395334__ihitokage__woosh-2.ogg"),       // jump sound
		rl.LoadSound("audio/490769__steveuk87__punch-2-heavy.ogg"), // punch sound
	}
	music := rl.LoadMusicStream("audio/378665__frankum__red-hard-lead-techno-house.mp3")

	blockers := make([]Blocker, 0)
	blockers = append(blockers, NewBlocker(-100, 600, 1480, 120, rl.Black)) // bottom of map
	blockers = append(blockers, NewBlocker(-100, -50, 1480, 50, rl.Black))  // top of map
	blockers = append(blockers, NewBlocker(-100, 0, 100, 720, rl.Black))    // left of map
	blockers = append(blockers, NewBlocker(1280, 0, 100, 720, rl.Black))    // right of map
	blockers = append(blockers, NewBlocker(200, 400, 200, 30, rl.Black))    // left platform
	blockers = append(blockers, NewBlocker(900, 400, 200, 30, rl.Black))    // right platform
	blockers = append(blockers, NewBlocker(550, 200, 200, 30, rl.Black))    // middle platform

	// create player one
	playerOne := NewCreature(rl.NewVector2(300, 520), SpriteSheets, rl.Green)

	// create player two
	playerTwo := NewCreature(rl.NewVector2(900, 520), SpriteSheets, rl.Red)
	playerTwo.Flip = -1

	// game over and rounds
	gameOver := false
	currentRound := 1

	// play music
	rl.PlayMusicStream(music)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Blue)

		if gameOver {
			rl.DrawText("GAME OVER", 500, 200, 30, rl.White)
			if playerOne.Victories < 2 && playerTwo.Victories < 2 { // draw condition
				rl.DrawText("NEITHER PLAYER WINS", 425, 230, 30, rl.White)
			} else if playerOne.Victories >= 2 { // player one win condition
				rl.DrawText("PLAYER ONE WINS", 450, 230, 30, rl.White)
			} else if playerTwo.Victories >= 2 { // player two win condition
				rl.DrawText("PLAYER TWO WINS", 450, 230, 30, rl.White)
			}
			rl.DrawText("PRESS 'R' TO START ANOTHER MATCH", 300, 260, 30, rl.White)
			if rl.IsKeyPressed(rl.KeyR) {
				gameOver = false
				// reset player one
				playerOne.Pos = rl.NewVector2(300, 520)
				playerOne.Flip = 1
				playerOne.Health = 100
				playerOne.Victories = 0
				// reset player two
				playerTwo.Pos = rl.NewVector2(900, 520)
				playerTwo.Flip = -1
				playerTwo.Health = 100
				playerTwo.Victories = 0
				// reset round
				currentRound = 1
			}
		} else {
			// handle player input
			HandlePlayerOne(&playerOne, playerOne.CanJump(blockers), SoundEffects)
			HandlePlayerTwo(&playerTwo, playerTwo.CanJump(blockers), SoundEffects)

			// check collision with blockers
			for i := 0; i < len(blockers); i++ {
				blockers[i].DrawBlocker()
				CheckCollision(&playerOne.PhysicsBody, blockers[i])
				CheckCollision(&playerTwo.PhysicsBody, blockers[i])
			}

			// check if players hit one another
			CheckPlayerHit(&playerOne, &playerTwo)

			// player 1 health bar
			rl.DrawRectangle(0, 0, 500, 100, rl.White)
			rl.DrawRectangle(10, 10, int32(480*(float32(playerOne.Health)/100)), 80, rl.Red)

			// player 2 health bar
			rl.DrawRectangle(780, 0, 500, 100, rl.White)
			temp1 := int32(480 * (float32(playerTwo.Health) / 100))
			temp2 := 480 - temp1
			rl.DrawRectangle(790+temp2, 10, temp1, 80, rl.Red)

			// round counter
			rl.DrawText("R"+strconv.Itoa(currentRound), 590, 0, 100, rl.White)

			// check round win, loss, and draw
			if playerOne.Health <= 0 && playerTwo.Health <= 0 { // draw
				currentRound++
				// reset player one
				playerOne.Pos = rl.NewVector2(300, 520)
				playerOne.Flip = 1
				playerOne.Health = 100
				// reset player two
				playerTwo.Pos = rl.NewVector2(900, 520)
				playerTwo.Flip = -1
				playerTwo.Health = 100
			} else if playerOne.Health <= 0 {
				currentRound++
				playerTwo.Victories++
				// reset player one
				playerOne.Pos = rl.NewVector2(300, 520)
				playerOne.Flip = 1
				playerOne.Health = 100
				// reset player two
				playerTwo.Pos = rl.NewVector2(900, 520)
				playerTwo.Flip = -1
				playerTwo.Health = 100
			} else if playerTwo.Health <= 0 {
				currentRound++
				playerOne.Victories++
				// reset player one
				playerOne.Pos = rl.NewVector2(300, 520)
				playerOne.Flip = 1
				playerOne.Health = 100
				// reset player two
				playerTwo.Pos = rl.NewVector2(900, 520)
				playerTwo.Flip = -1
				playerTwo.Health = 100
			}

			// check game over
			if currentRound > 3 {
				gameOver = true
			} else if playerOne.Victories == 2 {
				gameOver = true
			} else if playerTwo.Victories == 2 {
				gameOver = true
			}
		}

		rl.EndDrawing()
	}
}

func CheckPlayerHit(playerOne *Creature, playerTwo *Creature) {
	// get player one action hit box
	var rec1 rl.Rectangle = rl.NewRectangle(0, 0, 0, 0)
	if playerOne.isPunching {
		rec1 = playerOne.GetPlayerPunchHitBox()
	} else if playerOne.isKicking {
		rec1 = playerOne.GetPlayerKickHitBox()
	} else if playerOne.isBlocking {
		rec1 = playerOne.GetPlayerBlockHitBox()
	}

	// get player two action hit box
	var rec2 rl.Rectangle = rl.NewRectangle(0, 0, 0, 0)
	if playerTwo.isPunching {
		rec2 = playerTwo.GetPlayerPunchHitBox()
	} else if playerTwo.isKicking {
		rec2 = playerTwo.GetPlayerKickHitBox()
	} else if playerTwo.isBlocking {
		rec2 = playerTwo.GetPlayerBlockHitBox()
	}

	// check if player blocked attack
	if (playerOne.isBlocking || playerTwo.isBlocking) && ((playerOne.Flip == -1 && playerTwo.Flip == 1) || (playerOne.Flip == 1 && playerTwo.Flip == -1)) {
		if rl.CheckCollisionRecs(rec1, rec2) {
			playerOne.HitTimer = 0.3
			playerTwo.HitTimer = 0.3
			return
		}
		if rl.CheckCollisionRecs(rec1, playerTwo.GetPlayerRectangle()) {
			playerTwo.HitTimer = 0.3
			return
		}
		if rl.CheckCollisionRecs(rec2, playerOne.GetPlayerRectangle()) {
			playerOne.HitTimer = 0.3
			return
		}
	}

	// check if player one hit two
	if rec1.X != 0 && rec1.Y != 0 && rec1.Width != 0 && rec1.Height != 0 {
		if rl.CheckCollisionRecs(rec1, playerTwo.GetPlayerRectangle()) && playerTwo.HitTimer <= 0 {
			playerTwo.Health -= 10
			playerTwo.HitTimer = 0.3
		}
	}

	// check if player two hit one
	if rec2.X != 0 && rec2.Y != 0 && rec2.Width != 0 && rec2.Height != 0 {
		if rl.CheckCollisionRecs(rec2, playerOne.GetPlayerRectangle()) && playerOne.HitTimer <= 0 {
			playerOne.Health -= 10
			playerOne.HitTimer = 0.3
		}
	}

	// set hit / invincibility timers to prevents multiple hits at once
	if playerOne.HitTimer > 0 {
		playerOne.HitTimer -= rl.GetFrameTime()
		if playerOne.HitTimer < 0 {
			playerOne.HitTimer = 0
		}
	}
	if playerTwo.HitTimer > 0 {
		playerTwo.HitTimer -= rl.GetFrameTime()
		if playerTwo.HitTimer < 0 {
			playerTwo.HitTimer = 0
		}
	}
}

func HandlePlayerOne(playerOne *Creature, canJump bool, SoundEffects []rl.Sound) {
	playerOne.ApplyGravity(rl.NewVector2(0, 600))

	dir := 0
	if rl.IsKeyDown(rl.KeyA) {
		dir += -1
	}
	if rl.IsKeyDown(rl.KeyD) {
		dir += 1
	}
	if canJump && rl.IsKeyPressed(rl.KeyW) {
		playerOne.isJumping = true
		rl.PlaySound(SoundEffects[1])
		playerOne.Jump()
	}

	if rl.IsKeyPressed(rl.KeyF) && !playerOne.isKicking && !playerOne.isBlocking {
		playerOne.isPunching = true
		rl.PlaySound(SoundEffects[2])
	} else if rl.IsKeyPressed(rl.KeyG) && !playerOne.isPunching && !playerOne.isBlocking {
		playerOne.isKicking = true
		rl.PlaySound(SoundEffects[0])
	} else if rl.IsKeyDown(rl.KeyH) && !playerOne.isPunching && !playerOne.isKicking {
		playerOne.isBlocking = true
		dir = 0
	}

	if playerOne.isPunching {
		playerOne.AnimationStateMachine.ChangeState(PUNCHSTATE)
	} else if playerOne.isKicking {
		playerOne.AnimationStateMachine.ChangeState(KICKSTATE)
	} else if playerOne.isBlocking {
		playerOne.AnimationStateMachine.ChangeState(BLOCKSTATE)
	} else if playerOne.isJumping {
		playerOne.AnimationStateMachine.ChangeState(JUMPSTATE)
	} else if dir != 0 && !playerOne.isJumping {
		playerOne.AnimationStateMachine.ChangeState(WALKSTATE)
	} else {
		playerOne.AnimationStateMachine.ChangeState(IDLESTATE)
	}
	// rl.DrawRectangleLines(int32(playerOne.Pos.X), int32(playerOne.Pos.Y), int32(playerOne.Scale.X), int32(playerOne.Scale.Y), rl.Orange)
	// rl.DrawRectangleRec(playerOne.GetPlayerRectangle(), rl.Orange)
	playerOne.MoveCreature(dir)
	playerOne.UpdateCreature()
}

func HandlePlayerTwo(playerTwo *Creature, canJump bool, SoundEffects []rl.Sound) {
	playerTwo.ApplyGravity(rl.NewVector2(0, 600))

	dir := 0
	if rl.IsKeyDown(rl.KeyLeft) {
		dir += -1
	}
	if rl.IsKeyDown(rl.KeyRight) {
		dir += 1
	}
	if canJump && rl.IsKeyPressed(rl.KeyUp) {
		playerTwo.isJumping = true
		rl.PlaySound(SoundEffects[1])
		playerTwo.Jump()
	}

	if rl.IsKeyPressed(rl.KeyKp1) && !playerTwo.isKicking && !playerTwo.isBlocking {
		playerTwo.isPunching = true
		rl.PlaySound(SoundEffects[2])
	} else if rl.IsKeyPressed(rl.KeyKp2) && !playerTwo.isPunching && !playerTwo.isBlocking {
		playerTwo.isKicking = true
		rl.PlaySound(SoundEffects[0])
	} else if rl.IsKeyDown(rl.KeyKp3) && !playerTwo.isPunching && !playerTwo.isKicking {
		playerTwo.isBlocking = true
		dir = 0
	}

	if playerTwo.isPunching {
		playerTwo.AnimationStateMachine.ChangeState(PUNCHSTATE)
	} else if playerTwo.isKicking {
		playerTwo.AnimationStateMachine.ChangeState(KICKSTATE)
	} else if playerTwo.isBlocking {
		playerTwo.AnimationStateMachine.ChangeState(BLOCKSTATE)
	} else if playerTwo.isJumping {
		playerTwo.AnimationStateMachine.ChangeState(JUMPSTATE)
	} else if dir != 0 && !playerTwo.isJumping {
		playerTwo.AnimationStateMachine.ChangeState(WALKSTATE)
	} else {
		playerTwo.AnimationStateMachine.ChangeState(IDLESTATE)
	}
	// rl.DrawRectangleLines(int32(playerTwo.Pos.X), int32(playerTwo.Pos.Y), int32(playerTwo.Scale.X), int32(playerTwo.Scale.Y), rl.Orange)
	// rl.DrawRectangleRec(playerTwo.GetPlayerRectangle(), rl.Orange)
	playerTwo.MoveCreature(dir)
	playerTwo.UpdateCreature()
}
