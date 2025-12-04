package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	position rl.Vector2
	velocity rl.Vector2
	radius   float32
	launched bool
}

// the "math" package referred to in rubric cannot be used
// as it is not the actual default package
// func lerp(from float32, to float32, point float32) float32 {
// 	return from + point*(to-from)
// }

func (b *Ball) VelocityTick() {
	//fmt.Println(b.velocity.Y)
	adjustedVel := rl.Vector2Scale(b.velocity, rl.GetFrameTime())
	b.position = rl.Vector2Add(b.position, adjustedVel)
}

func (b *Ball) CheckForWalls() {
	if b.position.X+b.radius >= 800 {
		b.position.X = 790
		b.velocity.X *= -1
	}
	if b.position.X-b.radius <= 0 {
		b.position.X = 10
		b.velocity.X *= -1
	}
	if b.position.Y-b.radius <= 0 {
		b.position.Y = 10
		b.velocity.Y *= -1
	}
}

func (b *Ball) CheckForPaddle(p Paddle, blocksHit int) {
	if (b.position.X-b.radius < p.position.X+100) && (b.position.X+b.radius > p.position.X) &&
		(b.position.Y+b.radius < p.position.Y+5) && (b.position.Y+b.radius > p.position.Y) &&
		(b.velocity.Y > 0) {

		// get points
		lp := p.position.X
		rp := p.position.X + 100
		var increment float32 = 11 + (1.0 / 9.0) // nine regions on paddle

		// get region
		var region int = 0
		for i := lp; i <= rp-increment; i += increment {
			region++
			if b.position.X <= i+increment && b.position.X >= i {
				break
			}
		}

		// bounce based on region
		switch region {
		case 1, 9:
			b.velocity.Y = -155
			b.velocity.X = adjustXVelocity(region, 1, 9, 100)
		case 2, 8:
			b.velocity.Y = -180
			b.velocity.X = adjustXVelocity(region, 2, 8, 75)
		case 3, 7:
			b.velocity.Y = -205
			b.velocity.X = adjustXVelocity(region, 3, 7, 50)
		case 4, 6:
			b.velocity.Y = -230
			if b.velocity.X > -1 && b.velocity.X < 1 { // if x == 0
				b.velocity.X = 0
			} else if b.velocity.X < 0 {
				b.velocity.X = -280
			} else if b.velocity.X > 0 {
				b.velocity.X = 280
			}
		case 5:
			b.velocity.Y = -255
			if b.velocity.X > -1 && b.velocity.X < 1 { // if x == 0
				b.velocity.X = 0
			} else if b.velocity.X < 0 {
				b.velocity.X = -255
			} else if b.velocity.X > 0 {
				b.velocity.X = 255
			}
		}

		if b.velocity.Y < 0 {
			b.velocity.Y -= 1.5 * float32(blocksHit)
		} else if b.velocity.Y > 0 {
			b.velocity.Y += 1.5 * float32(blocksHit)
		}
		if b.velocity.X < 0 {
			b.velocity.X -= 1.5 * float32(blocksHit)
		} else if b.velocity.X > 0 {
			b.velocity.X += 1.5 * float32(blocksHit)
		}
	}
}

func adjustXVelocity(region int, b1 int, b2 int, adjust float32) float32 {
	if region == b1 {
		return -255 - adjust
	} else if region == b2 {
		return 255 + adjust
	}
	return -1 // I don't like Go forcing me to include a return statement here
}

func (b *Ball) CheckForBlock(blocks *[3][10]Block, blocksHit int) int {
	for i := 0; i < 3; i++ {
		for k := 0; k < 10; k++ {
			if !blocks[i][k].visible {
				continue
			}
			// did ball hit bottom or top
			if ((b.position.X-b.radius > blocks[i][k].position.X) && (b.position.X <= blocks[i][k].position.X+62)) ||
				((b.position.X+b.radius < blocks[i][k].position.X+62) && (b.position.X >= blocks[i][k].position.X)) {
				if ((b.position.Y-b.radius < blocks[i][k].position.Y+62) && (b.position.Y-b.radius > blocks[i][k].position.Y)) ||
					((b.position.Y+b.radius > blocks[i][k].position.Y) && (b.position.Y+b.radius < blocks[i][k].position.Y+62)) {
					b.velocity.Y *= -1
					blocks[i][k].visible = false
					blocksHit++
					if b.velocity.Y < 0 {
						b.velocity.Y -= 1.5
					} else if b.velocity.Y > 0 {
						b.velocity.Y += 1.5
					}
					if b.velocity.X < 0 {
						b.velocity.X -= 1.5
					} else if b.velocity.X > 0 {
						b.velocity.X += 1.5
					}
				}
			}
			if !blocks[i][k].visible {
				continue
			}

			// did ball hit left or right
			if b.position.Y <= blocks[i][k].position.Y+62 && b.position.Y >= blocks[i][k].position.Y {
				if ((b.position.X-b.radius < blocks[i][k].position.X+62) && (b.position.X-b.radius > blocks[i][k].position.X)) ||
					((b.position.X+b.radius > blocks[i][k].position.X) && (b.position.X+b.radius < blocks[i][k].position.X+62)) {
					b.velocity.X *= -1
					blocks[i][k].visible = false
					blocksHit++
					if b.velocity.Y < 0 {
						b.velocity.Y -= 1.5
					} else if b.velocity.Y > 0 {
						b.velocity.Y += 1.5
					}
					if b.velocity.X < 0 {
						b.velocity.X -= 1.5
					} else if b.velocity.X > 0 {
						b.velocity.X += 1.5
					}
				}
			}
		}
	}
	return blocksHit
}
