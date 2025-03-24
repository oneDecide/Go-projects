package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	Position     rl.Vector2
	Velocity     rl.Vector2
	Active       bool
	PunchThrough bool
	Lifetime     float32
}

var (
	projectiles     []Projectile
	baseVelocity    float32 = 500
	upgradeVelocity float32 = 0
)

func initProjectiles() {
	projectiles = make([]Projectile, 0)
}

func fireProjectile(position rl.Vector2, direction rl.Vector2) {
	projectile := Projectile{
		Position: position,
		Velocity: rl.Vector2Scale(direction, baseVelocity+upgradeVelocity),
		Active:   true,
		Lifetime: 3.0,
	}
	projectiles = append(projectiles, projectile)
}

func updateProjectiles() {
	dt := rl.GetFrameTime()

	for i := range projectiles {
		if projectiles[i].Active {
			projectiles[i].Position.X += projectiles[i].Velocity.X * dt
			projectiles[i].Position.Y += projectiles[i].Velocity.Y * dt

			projectiles[i].Lifetime -= dt

			if projectiles[i].Lifetime <= 0 {
				projectiles[i].Active = false
			}
		}
	}
}

func drawProjectiles() {
	for _, p := range projectiles {
		if p.Active {
			rl.DrawCircleV(p.Position, 5, rl.Red)
		}
	}
}

func upgradeProjectileVelocity() {
	upgradeVelocity += 50
}

func upgradeProjectilePunchThrough() {
	for i := range projectiles {
		projectiles[i].PunchThrough = true
	}
}
