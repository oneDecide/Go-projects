package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player struct {
		Position   rl.Vector2
		BodyRadius float32
		score      int
	}
)

func initPlayer() {
	player.Position = rl.NewVector2(screenWidth/2, screenHeight/2)
	player.BodyRadius = 100
}

func drawPlayer() {
	bodyColor := color.RGBA{25, 255, 255, 100}
	rl.DrawCircleV(player.Position, player.BodyRadius, bodyColor)
}

func updateMovement() {
	mousepos := getWorldMousePosition()
	player.Position = mousepos
}

func handleInput() {
	if rl.IsKeyPressed(rl.KeySpace) || rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		spawnMine()
		score++
	}
}
