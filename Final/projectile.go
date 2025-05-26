package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Projectile struct {
	Position    rl.Vector2
	Direction   rl.Vector2
	Speed       float32
	Damage      int
	Instances   int
	TimeCreated float32
	Color       rl.Color
	IsEnemy     bool
}

func NewProjectile(pos, dir rl.Vector2, damageLevel int, isEnemy bool) *Projectile {
	color := rl.Orange
	if isEnemy {
		color = rl.Red
	}
	return &Projectile{
		Position:    pos,
		Direction:   rl.Vector2Normalize(dir),
		Speed:       5,
		Damage:      5,
		Instances:   damageLevel,
		TimeCreated: float32(rl.GetTime()),
		Color:       color,
		IsEnemy:     isEnemy,
	}
}

func (p *Projectile) Update(delta float32) {
	p.Position.X += p.Direction.X * p.Speed * delta * 60
	p.Position.Y += p.Direction.Y * p.Speed * delta * 60
}

func (p *Projectile) Draw() {
	rl.DrawCircleV(p.Position, 5, p.Color)
}

func (p *Projectile) ShouldExpire() bool {
	return float32(rl.GetTime())-p.TimeCreated > 7
}
