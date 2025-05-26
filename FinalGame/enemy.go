package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy interface {
	Update()
	Draw()
	LevelUp()
	Hit(damage int) bool
	GetPosition() rl.Vector2
	CollidesPlayer(p *Player) bool
}

// CircleEnemy — melee attacker
type CircleEnemy struct {
	Pos       rl.Vector2
	Speed     float32
	Health    int
	MaxHealth int
	target    *Player
}

func NewCircleEnemy(p *Player) *CircleEnemy {
	return &CircleEnemy{
		Pos:       randomSpawn(),
		Speed:     100,
		Health:    MeleeHP,
		MaxHealth: MeleeHP,
		target:    p,
	}
}

func (e *CircleEnemy) Update() {
	dir := rl.Vector2Subtract(e.target.Pos, e.Pos)
	if rl.Vector2Length(dir) > 0 {
		step := rl.Vector2Scale(rl.Vector2Normalize(dir), e.Speed*rl.GetFrameTime())
		e.Pos = rl.Vector2Add(e.Pos, step)
	}
}

func (e *CircleEnemy) Draw() {
	rl.DrawCircleV(e.Pos, 15, rl.Red)
	// Health bar
	barW := int32(30)
	pct := float32(e.Health) / float32(e.MaxHealth)
	x := int32(e.Pos.X) - barW/2
	y := int32(e.Pos.Y) - 25
	rl.DrawRectangle(x, y, barW, 4, rl.Red)
	rl.DrawRectangle(x, y, int32(float32(barW)*pct), 4, rl.Green)
}

func (e *CircleEnemy) LevelUp() {
	e.MaxHealth += 10
	e.Health = e.MaxHealth
	e.Speed += 20
}

func (e *CircleEnemy) Hit(damage int) bool {
	e.Health -= damage
	return e.Health <= 0
}

func (e *CircleEnemy) GetPosition() rl.Vector2 { return e.Pos }

func (e *CircleEnemy) CollidesPlayer(p *Player) bool {
	return rl.CheckCollisionCircles(e.Pos, 15, p.Pos, 20)
}

// ArrowEnemy — ranged attacker
type ArrowEnemy struct {
	Pos        rl.Vector2
	Speed      float32
	Health     int
	MaxHealth  int
	target     *Player
	shootTimer float32
}

func NewArrowEnemy(p *Player) *ArrowEnemy {
	return &ArrowEnemy{
		Pos:       randomSpawn(),
		Speed:     80,
		Health:    RangedHP,
		MaxHealth: RangedHP,
		target:    p,
	}
}

func (e *ArrowEnemy) Update() {
	dir := rl.Vector2Subtract(e.target.Pos, e.Pos)
	dist := rl.Vector2Length(dir)
	if dist > RangedRadius {
		step := rl.Vector2Scale(rl.Vector2Normalize(dir), e.Speed*rl.GetFrameTime())
		e.Pos = rl.Vector2Add(e.Pos, step)
	} else {
		perp := rl.Vector2{X: -dir.Y, Y: dir.X}
		step := rl.Vector2Scale(rl.Vector2Normalize(perp), e.Speed*rl.GetFrameTime())
		e.Pos = rl.Vector2Add(e.Pos, step)
	}
	e.shootTimer += rl.GetFrameTime()
	if e.shootTimer > 1.4 {
		e.shootTimer = 0
		dirN := rl.Vector2Normalize(dir)
		proj := NewProjectile(e.Pos, dirN, 5, OwnerEnemy)
		gameRef.projectiles = append(gameRef.projectiles, proj)
	}
}

func (e *ArrowEnemy) Draw() {
	pts := []rl.Vector2{
		{X: e.Pos.X, Y: e.Pos.Y - 10},
		{X: e.Pos.X - 8, Y: e.Pos.Y + 10},
		{X: e.Pos.X + 8, Y: e.Pos.Y + 10},
	}
	rl.DrawTriangle(pts[0], pts[1], pts[2], rl.DarkPurple)
	// Health bar
	barW := int32(30)
	pct := float32(e.Health) / float32(e.MaxHealth)
	x := int32(e.Pos.X) - barW/2
	y := int32(e.Pos.Y) - 25
	rl.DrawRectangle(x, y, barW, 4, rl.Red)
	rl.DrawRectangle(x, y, int32(float32(barW)*pct), 4, rl.Green)
}

func (e *ArrowEnemy) LevelUp() {
	e.MaxHealth += 5
	e.Health = e.MaxHealth
}

func (e *ArrowEnemy) Hit(damage int) bool {
	e.Health -= damage
	return e.Health <= 0
}

func (e *ArrowEnemy) GetPosition() rl.Vector2 { return e.Pos }

func (e *ArrowEnemy) CollidesPlayer(p *Player) bool {
	return false
}
