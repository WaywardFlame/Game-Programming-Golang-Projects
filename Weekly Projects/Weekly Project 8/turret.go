package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Turret struct {
	Position     rl.Vector2
	Speed        int
	Angle        float32
	Texture      rl.Texture2D
	Lasers       []Projectile
	Firing       bool
	NumShots     int
	PreventStuck int
}

func newTurret(pos rl.Vector2, s int, a float32) Turret {
	tex := rl.LoadTexture("textures/playerTurret.png")
	lasers := make([]Projectile, 0, 1200)
	return Turret{Position: pos, Speed: s, Angle: a, Texture: tex, Lasers: lasers, Firing: false, PreventStuck: 0}
}

func (p *Turret) turretMove(tex1 rl.Texture2D, tex2 rl.Texture2D, sound rl.Sound, count *int) {
	// player rotate
	if !p.Firing {
		if rl.IsKeyDown(rl.KeyQ) {
			p.Angle -= float32(p.Speed) * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyE) {
			p.Angle += float32(p.Speed) * rl.GetFrameTime()
		}
		if rl.IsKeyPressed(rl.KeyW) {
			p.Lasers = append(p.Lasers, newProjectile(p.Position, float64(p.Angle), 300, tex1, tex2))
			p.Firing = true
			p.NumShots--
			*count++
			rl.PlaySound(sound)
		}
	}
}

func (p *Turret) addTrail(pb *Projectile, tex rl.Texture2D) {
	if pb.TrailTimer%120 == 0 {
		p.Lasers = append(p.Lasers, Projectile{Position: pb.Position, Angle: 0, Velocity: rl.Vector2Zero(), TextRect: tex})
	} else {
		pb.TrailTimer++
	}
}

func (p *Turret) checkBoundary(pb *Projectile) {
	if pb.Position.X > 580 || pb.Position.X < 20 || pb.Position.Y < 20 || pb.Position.Y > 880 {
		p.Firing = false
		p.Lasers = make([]Projectile, 0, 1200)
	}
}

func checkWallsAndMirrors(Walls []rl.Rectangle, Mirrors []rl.Rectangle, p *Turret) {
	// compare ball with g.Walls, stop laser if collision occurs
	lengthPlayer := len(p.Lasers)
	lengthWalls := len(Walls)
	if lengthPlayer > 0 && lengthWalls > 0 { // I don't really know why, but not including this crashes the game
		for i := 0; i < len(Walls); i++ {
			if p.Lasers[0].Position.X < Walls[i].X+Walls[i].Width && p.Lasers[0].Position.X > Walls[i].X {
				if p.Lasers[0].Position.Y < Walls[i].Y+Walls[i].Height && p.Lasers[0].Position.Y > Walls[i].Y {
					p.Firing = false
					p.Lasers = make([]Projectile, 0, 1200)
					return
				}
			}
		}
	}

	// I hate corners. The sides are no problem, but god forbid
	// the ball hits a corner.

	// compare ball with g.Mirrors, reflecting laser and change angle if collision occurs
	lengthMirrors := len(Mirrors)
	lengthPlayer = len(p.Lasers)
	if lengthPlayer > 0 && lengthMirrors > 0 && p.PreventStuck == 0 {
		pb := p.Lasers[0]
		for i := 0; i < len(Mirrors); i++ {
			// did ball hit bottom or top of mirror
			// 12 is radius of laser ball + 2
			if ((pb.Position.X-12 >= Mirrors[i].X) && (pb.Position.X <= Mirrors[i].X+Mirrors[i].Width)) ||
				((pb.Position.X+12 <= Mirrors[i].X+Mirrors[i].Width) && (pb.Position.X >= Mirrors[i].X)) {
				if ((pb.Position.Y-12 <= Mirrors[i].Y+Mirrors[i].Height) && (pb.Position.Y-12 >= Mirrors[i].Y)) ||
					((pb.Position.Y+12 >= Mirrors[i].Y) && (pb.Position.Y+12 <= Mirrors[i].Y+Mirrors[i].Height)) {
					p.Lasers[0].Velocity.Y *= -1
					p.PreventStuck = 10
					return
				}
			}
			// did ball hit left or right
			// 12 is radius of laser ball + 2
			if pb.Position.Y <= Mirrors[i].Y+Mirrors[i].Height && pb.Position.Y >= Mirrors[i].Y {
				if (pb.Position.X-12 <= Mirrors[i].X+Mirrors[i].Width) && (pb.Position.X-12 >= Mirrors[i].X) {
					if p.Lasers[0].Velocity.X < 0 {
						p.Lasers[0].Velocity.X *= -1
					}
					p.PreventStuck = 10
					return
				} else if (pb.Position.X+12 >= Mirrors[i].X) && (pb.Position.X+12 <= Mirrors[i].X+Mirrors[i].Width) {
					if p.Lasers[0].Velocity.X > 0 {
						p.Lasers[0].Velocity.X *= -1
					}
					p.PreventStuck = 10
					return
				}
			}
		}
	} else {
		p.PreventStuck--
		if p.PreventStuck < 0 {
			p.PreventStuck = 0
		}
	}
}
