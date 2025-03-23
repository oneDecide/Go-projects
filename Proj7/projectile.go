package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Projectile struct {
	Position     rl.Vector2
	Velocity     rl.Vector2
	Active       bool
	PunchThrough bool    // Can hit multiple enemies
	Lifetime     float32 // Time in seconds before the projectile is deactivated
}

var (
	projectiles     []Projectile
	baseVelocity    float32 = 500 // Adjusted for delta time
	upgradeVelocity float32 = 0
)

func initProjectiles() {
	projectiles = make([]Projectile, 0)
}

func fireProjectile(position rl.Vector2, direction rl.Vector2) {
	projectile := Projectile{
		Position:     position,
		Velocity:     rl.Vector2Scale(direction, baseVelocity+upgradeVelocity),
		Active:       true,
		PunchThrough: false, // Default to false, can be upgraded
		Lifetime:     3.0,   // Projectile lasts for 3 seconds
	}
	projectiles = append(projectiles, projectile)
}

func updateProjectiles() {
	dt := rl.GetFrameTime() // Get delta time

	for i := range projectiles {
		if projectiles[i].Active {
			// Update position
			projectiles[i].Position.X += projectiles[i].Velocity.X * dt
			projectiles[i].Position.Y += projectiles[i].Velocity.Y * dt

			// Decrease lifetime
			projectiles[i].Lifetime -= dt

			// Deactivate projectiles when their lifetime expires
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

// Upgrade functions
func upgradeProjectileVelocity() {
	upgradeVelocity += 50 // Adjusted for delta time
}

func upgradeProjectilePunchThrough() {
	for i := range projectiles {
		projectiles[i].PunchThrough = true
	}
}
