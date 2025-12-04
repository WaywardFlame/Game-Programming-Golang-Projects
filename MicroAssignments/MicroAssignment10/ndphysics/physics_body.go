package ndphysics

import rl "github.com/gen2brain/raylib-go/raylib"

type PhysicsBody struct {
	pos rl.Vector2
	vel rl.Vector2
}

func NewPhysicsBody(newPos rl.Vector2, newVel rl.Vector2) PhysicsBody {
	pb := PhysicsBody{pos: newPos, vel: newVel}
	return pb
}

func (pb *PhysicsBody) PhysicsUpdate() {
	pb.VelocityTick()
	//other stuff may be called here later
}

func (pb *PhysicsBody) VelocityTick() {
	adjustedVel := rl.Vector2Scale(pb.vel, rl.GetFrameTime())
	pb.pos = rl.Vector2Add(pb.pos, adjustedVel)
}

func (pb *PhysicsBody) GetVelocity() rl.Vector2 {
	return pb.vel
}

func (pb *PhysicsBody) GetPosition() rl.Vector2 {
	return pb.pos
}

func (pb *PhysicsBody) SetPosition(position rl.Vector2) {
	pb.pos = position
}
