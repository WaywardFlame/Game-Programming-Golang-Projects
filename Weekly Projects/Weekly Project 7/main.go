package main

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1280, 720, "Weekly Project 7")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)

	// load in textures
	textures := [6]rl.Texture2D{
		rl.LoadTexture("textures/LargeAsteroid.png"),
		rl.LoadTexture("textures/MediumAsteroid.png"),
		rl.LoadTexture("textures/SmallAsteroid.png"),
		rl.LoadTexture("textures/Planet.png"),
		rl.LoadTexture("textures/PlayerShip.png"),
		rl.LoadTexture("textures/PlayerBullet.png"),
	}

	// load in sound effects
	sounds := [4]rl.Sound{
		rl.LoadSound("audio/346116__lulyc__retro-game-heal-sound.wav"),
		rl.LoadSound("audio/391660__jeckkech__projectile.wav"),
		rl.LoadSound("audio/442127__euphrosyyn__8-bit-game-over"),
		rl.LoadSound("audio/450614__breviceps__8-bit-collect-sound-timer-countdown.wav"),
	}

	// create alias for explosions
	explosion := rl.LoadSound("audio/334266__aceofspadesproduc100__short-explosion-1.wav")
	rl.SetSoundVolume(explosion, 0.25)
	var expIndex int = 0
	explosions := make([]rl.Sound, 0, 10)
	for i := 0; i < 10; i++ {
		explosions = append(explosions, rl.LoadSoundAlias(explosion))
	}

	// create music stream
	music := rl.LoadMusicStream("audio/720970__universfield__my-dreams.mp3")

	// create planet
	planet := createPlanet(textures[3])

	// create player
	player := createPlayer(textures[4])

	// create projectile slice
	projectiles := make([]PhysicsBody, 0, 100)

	// create asteroids slice
	asteroids := make([]Asteroid, 0, 100)

	// per second variables
	// don't like having something framerate dependent
	// but it doesn't matter for this project
	var k int = 0

	// game over variable
	var gameOver bool = false

	// start music stream
	rl.PlayMusicStream(music)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(music)
		rl.BeginDrawing()

		if gameOver {
			rl.ClearBackground(rl.Red)
			rl.DrawText("Game Over - Press R to restart", 320, 320, 40, rl.White)
			if rl.IsKeyPressed(rl.KeyR) {
				planet = createPlanet(textures[3])
				player = createPlayer(textures[4])
				projectiles = make([]PhysicsBody, 0, 100)
				asteroids = make([]Asteroid, 0, 100)
				k = 0
				gameOver = false
			}
			if gameOver {
				rl.EndDrawing()
				continue
			}
		}

		rl.ClearBackground(rl.Black)
		k++

		// planet
		DrawBody(planet.texture, planet.position, 0, 1, rl.White)

		// player
		DrawBody(player.texture, player.position, player.angle, 1, rl.White)
		player.playerMove()

		// projectiles
		if rl.IsKeyPressed(rl.KeySpace) {
			projectiles = append(projectiles, newProjectile(player.position, float64(player.angle), float32(player.speed), textures[5]))
			rl.PlaySound(sounds[1]) // projectile sound
		}
		for i := 0; i < len(projectiles); i++ {
			projectiles[i].ProjectileUpdate()
			if projectiles[i].collide {
				DrawBody(projectiles[i].texture, projectiles[i].position, 0, 1, rl.White)
			}
		}

		// asteroids
		if k >= 60 { // spawn asteroid every two seconds
			asteroids = append(asteroids, createAsteroid(textures[0], 64, 3))
			k = 0
		}
		for i := 0; i < len(asteroids); i++ {
			asteroids[i].asteroidUpdate(&planet)
			if asteroids[i].collide {
				DrawBody(asteroids[i].texture, asteroids[i].position, 0, 1, asteroids[i].color)
			}
		}
		if planet.health <= 0 {
			gameOver = true
			rl.PlaySound(sounds[2])
			rl.EndDrawing()
			continue
		}

		// check collision between asteroids and projectiles
		for i := 0; i < len(projectiles); i++ {
			for j := 0; j < len(asteroids); j++ {
				if projectiles[i].collide && asteroids[j].collide {
					if rl.Vector2Distance(projectiles[i].position, asteroids[j].position) <= projectiles[i].radius+asteroids[j].radius {
						projectiles[i].collide = false
						asteroids[j].collide = false
						asteroids[j].breakup(&asteroids, textures)
						rl.PlaySound(explosions[expIndex])
						expIndex++
						if expIndex >= len(explosions) {
							expIndex = 0
						}
					}
				}
			}
		}

		// check player planet collision
		if rl.Vector2Distance(player.position, planet.position) <= player.radius+planet.radius {
			if player.cargo > 0 {
				planet.health += player.cargo * 5
				if planet.health > 100 {
					planet.health = 100
				}
				player.cargo = 0
				rl.PlaySound(sounds[0]) // planet heal
			}
		}

		// check player small asteroid collision
		for i := 0; i < len(asteroids); i++ {
			if asteroids[i].size == 1 && asteroids[i].collide {
				if rl.Vector2Distance(player.position, asteroids[i].position) <= player.radius+asteroids[i].radius {
					asteroids[i].collide = false
					player.cargo++
					rl.PlaySound(sounds[3])
				}
			}
		}

		// text
		rl.DrawText("Planet Health: "+strconv.Itoa(planet.health), 5, 5, 20, rl.White)
		rl.DrawText("Cargo: "+strconv.Itoa(player.cargo), 5, 25, 20, rl.White)

		rl.EndDrawing()
	}
}
