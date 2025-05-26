package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 450
	gridSize     = 50
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "The Sides of Shape")
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	game := NewGame()
	defer game.Cleanup()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		game.Update()
		game.Draw()
	}
}
