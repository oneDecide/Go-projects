package ndphysics

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	PhysicsBody
	size float32
}

func NewProjectile(newSize float32, newPos rl.Vector2, newVel rl.Vector2) Projectile {
	pb := NewPhysicsBody(newPos, newVel, newSize*1.25)
	nc := Projectile{size: newSize, PhysicsBody: pb}

	return nc
}

func (p Projectile) DrawProjectile() {
	rl.DrawCircle(int32(p.Pos.X), int32(p.Pos.Y), p.size, rl.White)
	p.DrawBoundary()
}
