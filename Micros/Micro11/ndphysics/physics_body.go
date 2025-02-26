package ndphysics

import rl "github.com/gen2brain/raylib-go/raylib"

type PhysicsBody struct {
	Pos              rl.Vector2
	Vel              rl.Vector2
	Radius           float32
	Gravity          float32
	ignoreCollisions bool
}

func NewPhysicsBody(newPos rl.Vector2, newVel rl.Vector2, newRadius float32) PhysicsBody {
	pb := PhysicsBody{Pos: newPos, Vel: newVel, Radius: newRadius, Gravity: 0}
	pb.ignoreCollisions = false
	return pb
}

func (pb *PhysicsBody) CheckIntersection(otherPb *PhysicsBody) bool {
	if rl.Vector2Distance(pb.Pos, otherPb.Pos) <= pb.Radius+otherPb.Radius {
		pb.Bounce()
		return true
	}
	return false
}

func (pb *PhysicsBody) Bounce() {
	pb.Vel = rl.Vector2Scale(pb.Vel, -1)
}

func (pb PhysicsBody) DrawBoundary() {
	rl.DrawCircleLines(int32(pb.Pos.X), int32(pb.Pos.Y), pb.Radius, rl.Lime)
}

func (pb *PhysicsBody) SetIgnoreCollisions(ignore bool) {
	pb.ignoreCollisions = ignore
}

func (pb *PhysicsBody) PhysicsUpdate() {
	pb.GravityTick()
	pb.VelocityTick()
	//other stuff may be called here later
}

func (pb *PhysicsBody) VelocityTick() {
	adjustedVel := rl.Vector2Scale(pb.Vel, rl.GetFrameTime())
	pb.Pos = rl.Vector2Add(pb.Pos, adjustedVel)
}

func (pb *PhysicsBody) GravityTick() {
	pb.Vel = rl.Vector2Add(pb.Vel, rl.Vector2Scale(rl.NewVector2(0, pb.Gravity), rl.GetFrameTime()))
}
