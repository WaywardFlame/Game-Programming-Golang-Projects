package ndphysics

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsBody struct {
	pos              rl.Vector2
	vel              rl.Vector2
	radius           float32
	ignoreCollisions bool
}

func NewPhysicsBody(newPos rl.Vector2, newVel rl.Vector2, newRadius float32) PhysicsBody {
	pb := PhysicsBody{pos: newPos, vel: newVel, radius: newRadius}
	pb.ignoreCollisions = false
	return pb
}

func (pb *PhysicsBody) CheckIntersection(otherPb *PhysicsBody) bool {
	if rl.Vector2Distance(pb.pos, otherPb.pos) <= pb.radius+otherPb.radius {
		pb.Bounce(otherPb)
		return true
	}
	return false
}

func (pb *PhysicsBody) Bounce(otherPb *PhysicsBody) {
	// utilized a concept called "penetration resolution" to solve collisions
	distance := rl.NewVector2(pb.pos.X-otherPb.pos.X, pb.pos.Y-otherPb.pos.Y)
	magnitude := calculateMagnitude(distance)
	depth := pb.radius + otherPb.radius - magnitude
	res := calculateRes(distance, magnitude, depth)
	pb.pos = rl.Vector2Add(pb.pos, res)
	otherPb.pos = rl.Vector2Add(otherPb.pos, rl.Vector2Scale(res, -1))

	// bounce off
	pb.vel = rl.Vector2Scale(pb.vel, -1)
	otherPb.vel = rl.Vector2Scale(pb.vel, -1)
}

func calculateMagnitude(distance rl.Vector2) float32 {
	xSqr := distance.X * distance.X
	ySqr := distance.Y * distance.Y
	return float32(math.Sqrt(float64(xSqr + ySqr)))
}

func calculateRes(distance rl.Vector2, magnitude float32, depth float32) rl.Vector2 {
	unit := rl.NewVector2(distance.X/magnitude, distance.Y/magnitude)
	unit.X *= depth / 2
	unit.Y *= depth / 2
	return unit
}

func (pb PhysicsBody) DrawBoundary() {
	rl.DrawCircleLines(int32(pb.pos.X), int32(pb.pos.Y), pb.radius, rl.Lime)
}

func (pb *PhysicsBody) SetIgnoreCollisions(ignore bool) {
	pb.ignoreCollisions = ignore
}

func (pb *PhysicsBody) PhysicsUpdate() {
	pb.VelocityTick()
	//other stuff may be called here later
}

func (pb *PhysicsBody) VelocityTick() {
	adjustedVel := rl.Vector2Scale(pb.vel, rl.GetFrameTime())
	pb.pos = rl.Vector2Add(pb.pos, adjustedVel)
}
