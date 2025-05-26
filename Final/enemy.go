package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Enemy struct {
	Position   rl.Vector2
	Health     int
	Speed      float32
	EnemyType  int 
	LastAttack float32
	IsActive   bool
	Color      rl.Color
}

func NewEnemy(pos rl.Vector2, enemyType int) *Enemy {
	color := rl.Green
	if enemyType == 1 {
		color = rl.Purple
	}
	return &Enemy{
		Position:  pos,
		Health:    10,
		Speed:     1.5,
		EnemyType: enemyType,
		IsActive:  true,
		Color:     color,
	}
}

func (e *Enemy) Update(target rl.Vector2, delta float32, game *Game) {
	switch e.EnemyType {
	case 0: // Melee
		direction := rl.Vector2Subtract(target, e.Position)
		direction = rl.Vector2Normalize(direction)
		e.Position.X += direction.X * e.Speed * delta * 60
		e.Position.Y += direction.Y * e.Speed * delta * 60

		if rl.Vector2Distance(e.Position, target) < 20 {
			if float32(rl.GetTime())-e.LastAttack > 1.0/6.0 {
				e.LastAttack = float32(rl.GetTime())
			}
		}

	case 1: // Ranged
		distance := rl.Vector2Distance(e.Position, target)
		direction := rl.Vector2Subtract(target, e.Position)
		direction = rl.Vector2Normalize(direction)

		if distance < 200 {
			e.Position.X -= direction.X * e.Speed * delta * 60
			e.Position.Y -= direction.Y * e.Speed * delta * 60
		} else if distance > 250 {
			e.Position.X += direction.X * e.Speed * delta * 60
			e.Position.Y += direction.Y * e.Speed * delta * 60
		}

		if float32(rl.GetTime())-e.LastAttack > 2.0 {
			game.EnemyProjectiles = append(game.EnemyProjectiles,
				NewProjectile(e.Position, direction, 1, true))
			e.LastAttack = float32(rl.GetTime())
		}
	}
}

func (e *Enemy) Draw() {
	rl.DrawCircleV(e.Position, 15, e.Color)
	rl.DrawRectangle(
		int32(e.Position.X-15),
		int32(e.Position.Y-25),
		int32(30*(float32(e.Health)/10)),
		5,
		rl.Red,
	)
}

func (e *Enemy) TakeDamage(amount int) {
	e.Health -= amount
	if e.Health <= 0 {
		e.IsActive = false 
	}
}
