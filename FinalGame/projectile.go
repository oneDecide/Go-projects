package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Projectile travels and deals damage.
type Projectile struct {
	Pos       rl.Vector2
	Dir       rl.Vector2
	Speed     float32
	Damage    int
	Instances int
	SpawnTime float64 // now matches rl.GetTime() (float64)
	Owner     int
}

func NewProjectile(pos, dir rl.Vector2, dmg, owner int) *Projectile {
	return &Projectile{
		Pos:       pos,
		Dir:       dir,
		Speed:     400,
		Damage:    dmg,
		Instances: 1,
		SpawnTime: rl.GetTime(),
		Owner:     owner,
	}
}

// Update moves the projectile and returns false when expired.
func (p *Projectile) Update() bool {
	dt := rl.GetFrameTime()
	p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(p.Dir, p.Speed*dt))
	// Compare two float64s now
	return rl.GetTime()-p.SpawnTime <= float64(ProjectileLifespan)
}

func (p *Projectile) Draw() {
	col := rl.Orange
	if p.Owner == OwnerEnemy {
		col = rl.Red
	}
	rl.DrawCircleV(p.Pos, 5, col)
}

func (p *Projectile) Collides(e Enemy) bool {
	return rl.CheckCollisionCircles(p.Pos, 5, e.GetPosition(), 15)
}

func (p *Projectile) CollidesPlayer(pl *Player) bool {
	return rl.CheckCollisionCircles(p.Pos, 5, pl.Pos, 20)
}

func (p *Projectile) DamageInstance() int   { return p.Damage }
func (p *Projectile) ConsumeInstance() bool { p.Instances--; return p.Instances <= 0 }
