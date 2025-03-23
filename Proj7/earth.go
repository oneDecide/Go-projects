package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	earth struct {
		Position rl.Vector2
		Radius   float32
		Health   int // Replace Points with Health
	}
)

func initEarth(x, y float32) {
	earth.Position = rl.NewVector2(x, y)
	earth.Radius = 40  // Double the size (previously 20)
	earth.Health = 100 // Initial health
}

func updateEarth() {
	// Check for collision with player
	if rl.CheckCollisionCircles(player.Position, player.BodyRadius, earth.Position, earth.Radius) {
		// Restore health based on player's collected bits
		earth.Health += player.CollectedBits * 10
		player.CollectedBits = 0
	}

	// Check for collision with asteroids
	for i := range asteroids {
		if asteroids[i].Active && asteroids[i].Layer > 1 && rl.CheckCollisionCircles(asteroids[i].Position, asteroids[i].Radius, earth.Position, earth.Radius) {
			// Reduce Earth's health by 5
			earth.Health -= 5
			asteroids[i].Active = false // Deactivate the asteroid
		}
	}

	// Ensure health doesn't go below 0
	if earth.Health < 0 {
		earth.Health = 0
	}
}

func drawEarth() {
	// Draw Earth (circle)
	rl.DrawCircleV(earth.Position, earth.Radius, rl.Green)
}

func drawUI() {
	// Draw health and bits in the top-left corner of the screen
	rl.DrawText(fmt.Sprintf("Health: %d", earth.Health), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Bits: %d", player.CollectedBits), 10, 40, 20, rl.White)
}
