package ndphysics

import rl "github.com/gen2brain/raylib-go/raylib"

var globalRestitution float32 = 1

type PhysicsBody struct {
	Pos              rl.Vector2
	Vel              rl.Vector2
	Mass             float32
	Radius           float32
	Gravity          float32
	ignoreCollisions bool
}

func NewPhysicsBody(newPos rl.Vector2, newVel rl.Vector2, mass float32, newRadius float32) PhysicsBody {
	pb := PhysicsBody{Pos: newPos, Vel: newVel, Radius: newRadius, Gravity: 0}
	pb.Mass = mass
	pb.ignoreCollisions = false
	return pb
}

func (pb *PhysicsBody) CheckIntersection(otherPb *PhysicsBody) bool {
	if rl.Vector2Distance(pb.Pos, otherPb.Pos) <= pb.Radius+otherPb.Radius {
		pb.Bounce(otherPb)
		return true
	}
	return false
}

//func (pb *PhysicsBody) Bounce() {
//	pb.Vel = rl.Vector2Scale(pb.Vel, -1)
//}

func (pb *PhysicsBody) Bounce(otherPb *PhysicsBody) {

	normal := rl.Vector2Normalize(rl.Vector2Subtract(pb.Pos, otherPb.Pos))

	relativeVel := rl.Vector2Subtract(pb.Vel, otherPb.Vel)

	speedAlongNormal := rl.Vector2DotProduct(relativeVel, normal)

	if speedAlongNormal > 0 {
		return
	}

	impulseScalar := -(1 + globalRestitution) * speedAlongNormal
	impulseScalar /= (1 / pb.Mass) + (1 / otherPb.Mass)

	impulse := rl.Vector2Scale(normal, impulseScalar)
	pb.Vel = rl.Vector2Add(pb.Vel, rl.Vector2Scale(impulse, 1/pb.Mass))
	otherPb.Vel = rl.Vector2Subtract(otherPb.Vel, rl.Vector2Scale(impulse, 1/otherPb.Mass))

}

/*func (pb *PhysicsBody) Bounce(otherpb *PhysicsBody) {
	pb.Vel = rl.Vector2Scale(
		rl.Vector2Normalize(
			rl.Vector2Subtract(pb.Pos, otherpb.Pos)),
		rl.Vector2Length(pb.Vel),
	)
	//teleport the particle away immediately!
	teleportDir := rl.Vector2Normalize(rl.Vector2Subtract(pb.Pos, otherpb.Pos))
	teleportMag := pb.Radius + otherpb.Radius - rl.Vector2Distance(pb.Pos, otherpb.Pos)
	teleportFinal := rl.Vector2Scale(teleportDir, teleportMag)
	pb.Pos = rl.Vector2Add(pb.Pos, teleportFinal)
}*/

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
