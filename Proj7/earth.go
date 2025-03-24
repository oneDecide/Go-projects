package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	earth struct {
		Position rl.Vector2
		Radius   float32
		Health   int
	}
)

func initEarth(x, y float32) {
	earth.Position = rl.NewVector2(x, y)
	earth.Radius = 40
	earth.Health = 100
}

func updateEarth() {
	if rl.CheckCollisionCircles(player.Position, player.BodyRadius, earth.Position, earth.Radius) {
		if player.CollectedBits > 0 {
			earth.Health += player.CollectedBits * 10
			player.CollectedBits = 0
			playSoundWithPitch(depositSound)
		}
	}

	for i := range asteroids {
		if asteroids[i].Active && asteroids[i].Layer > 1 && rl.CheckCollisionCircles(asteroids[i].Position, asteroids[i].Radius, earth.Position, earth.Radius) {
			
			earth.Health -= 5
			asteroids[i].Active = false
		}
	}

	if earth.Health < 0 {
		earth.Health = 0
	}
}

func drawEarth() {
	rl.DrawCircleV(earth.Position, earth.Radius, rl.Green)
}

func drawUI() {
	rl.DrawText(fmt.Sprintf("Health: %d", earth.Health), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Bits: %d/2", player.CollectedBits), 10, 40, 20, rl.White)
}
