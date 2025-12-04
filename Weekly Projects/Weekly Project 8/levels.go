package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Originally planned for saving and loading
// hence the GameData struct. I kept it for the
// sake of readily creating data across levels.
type GameData struct {
	Level       int
	Player      Turret
	Walls       []rl.Rectangle
	Mirrors     []rl.Rectangle
	RedBlocks   []rl.Rectangle
	GreenBlock  rl.Rectangle
	Textures    []rl.Texture2D
	Sounds      []rl.Sound
	Initialized bool
	LasersFired int
	GameOvers   int
}

func DrawBody(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)
	rl.DrawTexturePro(texture, sourceRect, destRect, origin, angle, color)
}

func (g *GameData) drawPlayerAndLasers() {
	g.Player.turretMove(g.Textures[0], g.Textures[1], g.Sounds[1], &g.LasersFired)
	if g.Player.Firing {
		g.Player.addTrail(&g.Player.Lasers[0], g.Textures[1])
		g.Player.checkBoundary(&g.Player.Lasers[0])
	}
	for i := 0; i < len(g.Player.Lasers); i++ {
		if i == 0 {
			DrawBody(g.Player.Lasers[i].TextBall, g.Player.Lasers[i].Position, float32(g.Player.Lasers[i].Angle), 1, rl.White)
		} else {
			DrawBody(g.Player.Lasers[i].TextRect, g.Player.Lasers[i].Position, float32(g.Player.Lasers[i].Angle), 1, rl.White)
		}
		g.Player.Lasers[i].ProjectileUpdate()
	}
	DrawBody(g.Player.Texture, g.Player.Position, g.Player.Angle, 1, rl.White)
}

func drawWallsAndMirrors(Walls []rl.Rectangle, Mirrors []rl.Rectangle) {
	for i := 0; i < len(Walls); i++ {
		rl.DrawRectangle(int32(Walls[i].X), int32(Walls[i].Y), int32(Walls[i].Width), int32(Walls[i].Height), rl.White)
	}
	for i := 0; i < len(Mirrors); i++ {
		rl.DrawRectangle(int32(Mirrors[i].X), int32(Mirrors[i].Y), int32(Mirrors[i].Width), int32(Mirrors[i].Height), rl.SkyBlue)
	}
}

func (g *GameData) drawMainMenu() {
	if !g.Initialized {
		g.LasersFired = 0
		g.GameOvers = 0
		g.Player.Firing = false
		g.Player.Lasers = make([]Projectile, 0, 1000)
		g.Player.Angle = 0
		g.Walls = make([]rl.Rectangle, 0, 10)
		g.Mirrors = make([]rl.Rectangle, 0, 10)

		g.Walls = append(g.Walls, rl.NewRectangle(290, 450, 20, 100))

		g.Mirrors = append(g.Mirrors, rl.NewRectangle(100, 450, 100, 20))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(400, 450, 100, 20))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(30, 650, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 650, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(100, 850, 100, 20))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(400, 850, 100, 20))

		g.Initialized = true
	}
	g.drawPlayerAndLasers()
	drawWallsAndMirrors(g.Walls, g.Mirrors)
	checkWallsAndMirrors(g.Walls, g.Mirrors, &g.Player) // in turret
}

func (g *GameData) initializeLevelOne() {
	if !g.Initialized {
		g.Level = 1
		g.Player.Firing = false
		g.Player.Lasers = make([]Projectile, 0, 1000)
		g.Player.Angle = 0
		g.Player.NumShots = 5

		g.Walls = make([]rl.Rectangle, 0, 10)
		g.Mirrors = make([]rl.Rectangle, 0, 10)

		g.Walls = append(g.Walls, rl.NewRectangle(225, 300, 150, 20))

		// left path
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(50, 550, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(230, 400, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(50, 250, 20, 100))

		// right path
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 550, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(350, 400, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 250, 20, 100))

		// green block
		g.GreenBlock = rl.NewRectangle(250, 100, 100, 100)

		g.Initialized = true
	}
	g.drawPlayerAndLasers()
	drawWallsAndMirrors(g.Walls, g.Mirrors)
	checkWallsAndMirrors(g.Walls, g.Mirrors, &g.Player) // in turret
	rl.DrawRectangle(int32(g.GreenBlock.X), int32(g.GreenBlock.Y), int32(g.GreenBlock.Width), int32(g.GreenBlock.Height), rl.Green)
	g.checkGreenBlock()
	if g.Player.NumShots <= 0 && len(g.Player.Lasers) <= 0 {
		rl.PlaySound(g.Sounds[0])
		g.Initialized = false
		g.GameOvers++
	}
}

func (g *GameData) initializeLevelTwo() {
	if !g.Initialized {
		g.Player.Firing = false
		g.Player.Lasers = make([]Projectile, 0, 1000)
		g.Player.Angle = 0
		g.Player.NumShots = 5

		g.Walls = make([]rl.Rectangle, 0, 10)
		g.Mirrors = make([]rl.Rectangle, 0, 10)
		g.RedBlocks = make([]rl.Rectangle, 0, 5)

		g.GreenBlock = rl.NewRectangle(250, 50, 100, 100)

		g.Walls = append(g.Walls, rl.NewRectangle(30, 50, 200, 20))
		g.Walls = append(g.Walls, rl.NewRectangle(370, 50, 200, 20))
		g.Walls = append(g.Walls, rl.NewRectangle(200, 270, 20, 400))
		g.Walls = append(g.Walls, rl.NewRectangle(380, 270, 20, 400))
		g.Walls = append(g.Walls, rl.NewRectangle(200, 250, 200, 20))

		// left path
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(30, 600, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(170, 550, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(30, 450, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(170, 400, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(30, 300, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(170, 250, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(30, 150, 20, 100))

		// right path
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 600, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 450, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 300, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 150, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(410, 550, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(410, 400, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(410, 250, 20, 100))

		// red blocks
		g.RedBlocks = append(g.RedBlocks, rl.NewRectangle(450, 150, 50, 50))
		g.RedBlocks = append(g.RedBlocks, rl.NewRectangle(100, 150, 50, 50))

		g.Initialized = true
	}
	g.drawPlayerAndLasers()
	drawWallsAndMirrors(g.Walls, g.Mirrors)
	checkWallsAndMirrors(g.Walls, g.Mirrors, &g.Player) // in turret
	rl.DrawRectangle(int32(g.GreenBlock.X), int32(g.GreenBlock.Y), int32(g.GreenBlock.Width), int32(g.GreenBlock.Height), rl.Green)
	for i := 0; i < len(g.RedBlocks); i++ {
		rl.DrawRectangle(int32(g.RedBlocks[i].X), int32(g.RedBlocks[i].Y), int32(g.RedBlocks[i].Width), int32(g.RedBlocks[i].Height), rl.Red)
	}
	g.checkGreenBlock()
	g.checkRedBlocks()
	if g.Player.NumShots <= 0 && len(g.Player.Lasers) <= 0 {
		rl.PlaySound(g.Sounds[0])
		g.Initialized = false
		g.Level = 1
		g.GameOvers++
	}
}

func (g *GameData) initializeLevelThree() {
	if !g.Initialized {
		g.Player.Firing = false
		g.Player.Lasers = make([]Projectile, 0, 1000)
		g.Player.Angle = 0
		g.Player.NumShots = 5

		g.Walls = make([]rl.Rectangle, 0, 10)
		g.Mirrors = make([]rl.Rectangle, 0, 10)
		g.RedBlocks = make([]rl.Rectangle, 0, 5)

		g.GreenBlock = rl.NewRectangle(350, 100, 100, 100)

		g.Walls = append(g.Walls, rl.NewRectangle(30, 550, 400, 20))
		g.Walls = append(g.Walls, rl.NewRectangle(170, 300, 400, 20))

		g.Mirrors = append(g.Mirrors, rl.NewRectangle(550, 500, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(350, 330, 100, 20))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(200, 520, 100, 20))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(30, 250, 20, 100))
		g.Mirrors = append(g.Mirrors, rl.NewRectangle(200, 30, 100, 20))

		g.RedBlocks = append(g.RedBlocks, rl.NewRectangle(150, 150, 50, 50))

		g.Initialized = true
	}
	g.drawPlayerAndLasers()
	drawWallsAndMirrors(g.Walls, g.Mirrors)
	checkWallsAndMirrors(g.Walls, g.Mirrors, &g.Player) // in turret
	rl.DrawRectangle(int32(g.GreenBlock.X), int32(g.GreenBlock.Y), int32(g.GreenBlock.Width), int32(g.GreenBlock.Height), rl.Green)
	for i := 0; i < len(g.RedBlocks); i++ {
		rl.DrawRectangle(int32(g.RedBlocks[i].X), int32(g.RedBlocks[i].Y), int32(g.RedBlocks[i].Width), int32(g.RedBlocks[i].Height), rl.Red)
	}
	g.checkGreenBlock()
	g.checkRedBlocks()
	if g.Player.NumShots <= 0 && len(g.Player.Lasers) <= 0 {
		rl.PlaySound(g.Sounds[0])
		g.Initialized = false
		g.Level = 1
		g.GameOvers++
	}
}

func (g *GameData) checkGreenBlock() {
	if len(g.Player.Lasers) > 0 {
		if g.Player.Lasers[0].Position.X-8 < g.GreenBlock.X+g.GreenBlock.Width && g.Player.Lasers[0].Position.X+8 > g.GreenBlock.X {
			if g.Player.Lasers[0].Position.Y-8 < g.GreenBlock.Y+g.GreenBlock.Height && g.Player.Lasers[0].Position.Y+8 > g.GreenBlock.Y {
				g.Level++
				g.Initialized = false
				rl.PlaySound(g.Sounds[2])
			}
		}
	}
}

func (g *GameData) checkRedBlocks() {
	if len(g.Player.Lasers) > 0 && len(g.RedBlocks) > 0 { // I hate doing this
		for i := 0; i < len(g.RedBlocks); i++ {
			if g.Player.Lasers[0].Position.X-8 < g.RedBlocks[i].X+g.RedBlocks[i].Width && g.Player.Lasers[0].Position.X+8 > g.RedBlocks[i].X {
				if g.Player.Lasers[0].Position.Y-8 < g.RedBlocks[i].Y+g.RedBlocks[i].Height && g.Player.Lasers[0].Position.Y+8 > g.RedBlocks[i].Y {
					g.Player.Firing = false
					g.Player.Lasers = make([]Projectile, 0, 1200)
					g.RedBlocks = append(g.RedBlocks[:i], g.RedBlocks[i+1:]...)
					return
				}
			}
		}
	}
}
