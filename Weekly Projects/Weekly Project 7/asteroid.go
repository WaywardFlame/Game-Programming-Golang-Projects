package main

import (
	"math"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Asteroid struct {
	PhysicsBody
	color  rl.Color // light grey - dark grey - red
	size   int      // big - medium - small
	health int
}

// randomly choose which side of screen to spawn from
// then choose where along that side to spawn
func createAsteroid(texture rl.Texture2D, radius float32, bms int) Asteroid {
	side := rand.IntN(4)
	var where int
	pos := rl.Vector2Zero()
	if side == 0 { // top
		where = rand.IntN(1280)
		pos.X = float32(where)
		pos.Y = -64
	} else if side == 1 { // right
		where = rand.IntN(720)
		pos.Y = float32(where)
		pos.X = 1280 + 64
	} else if side == 2 { // bottom
		where = rand.IntN(1280)
		pos.X = float32(where)
		pos.Y = 720 + 64
	} else if side == 3 { // left
		where = rand.IntN(720)
		pos.Y = float32(where)
		pos.X = -64
	}

	// calculate direction and get normalized velocity - towards 640, 360
	dx := 640 - pos.X // target - current
	dy := 360 - pos.Y // target - current
	v := math.Sqrt(float64(dx*dx) + float64(dy*dy))
	dx = (dx / float32(v)) * 150
	dy = (dy / float32(v)) * 150
	vel := rl.NewVector2(dx, dy)

	// create physics body
	pb := newPhysicsBody(pos, vel, radius, texture, true)

	// create slightly random color
	r := uint8(rand.IntN(40)) + 216
	g := uint8(rand.IntN(40)) + 216
	b := uint8(rand.IntN(40)) + 216
	col := rl.NewColor(r, g, b, 255)

	return Asteroid{PhysicsBody: pb, color: col, size: bms, health: 100}
}

func (a *Asteroid) asteroidUpdate(planet *Planet) {
	// check planet collision
	if a.collide {
		a.checkPlanetCollision(planet)
	}

	// update asteroid position
	a.position = rl.Vector2Add(a.position, rl.Vector2Scale(a.velocity, rl.GetFrameTime()))
}

func (a *Asteroid) checkPlanetCollision(p *Planet) {
	if rl.Vector2Distance(a.position, p.position) <= a.radius+p.radius+8 {
		if a.size == 3 {
			p.health -= 10
		} else if a.size == 2 {
			p.health -= 5
		}
		a.collide = false
	}
}

func (a *Asteroid) breakup(asteroids *[]Asteroid, textures [6]rl.Texture2D) {
	if a.size == 3 {
		// create new asteroids
		a1 := createAsteroid(textures[1], 32, 2)
		a2 := createAsteroid(textures[1], 32, 2)
		//adjust position of new asteroids
		a1.position = a.position
		a2.position = a.position
		//adjust velocity of new asteroids
		//a1
		por1 := positiveOrNegative()
		por2 := positiveOrNegative()
		a1.velocity.X = float32(rand.IntN(25)+51) * por1
		a1.velocity.Y = float32(rand.IntN(25)+51) * por2
		//a2
		por1 = positiveOrNegative()
		por2 = positiveOrNegative()
		a2.velocity.X = float32(rand.IntN(25)+51) * por1
		a2.velocity.Y = float32(rand.IntN(25)+51) * por2
		// add asteroids to slice
		*asteroids = append(*asteroids, a1)
		*asteroids = append(*asteroids, a2)
	} else if a.size == 2 {
		// create new asteroids
		a1 := createAsteroid(textures[2], 16, 1)
		a2 := createAsteroid(textures[2], 16, 1)
		// adjust position of new asteroids
		a1.position = a.position
		a2.position = a.position
		//adjust velocity of new asteroids
		//a1
		por1 := positiveOrNegative()
		por2 := positiveOrNegative()
		a1.velocity.X = float32(rand.IntN(10)+1) * por1
		a1.velocity.Y = float32(rand.IntN(10)+1) * por2
		//a2
		por1 = positiveOrNegative()
		por2 = positiveOrNegative()
		a2.velocity.X = float32(rand.IntN(10)+1) * por1
		a2.velocity.Y = float32(rand.IntN(10)+1) * por2
		// add asteroids to slice
		*asteroids = append(*asteroids, a1)
		*asteroids = append(*asteroids, a2)
	}
}

func positiveOrNegative() float32 {
	if rand.IntN(2) == 0 {
		return 1
	} else {
		return -1
	}
}
