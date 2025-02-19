package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerColor := rl.NewColor(128, 200, 255, 255)
	playerSprite := rl.LoadTexture("textures/cyclops.png")

	var playerCreature Creature = NewCreature(rl.NewVector2(50, 50), 200, 50, playerSprite, playerColor)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		tempMove := rl.Vector2Zero()
		if rl.IsKeyDown(rl.KeyA) {
			tempMove.X -= 1
		}
		if rl.IsKeyDown(rl.KeyD) {
			tempMove.X += 1
		}
		if rl.IsKeyDown(rl.KeyW) {
			tempMove.Y -= 1
		}
		if rl.IsKeyDown(rl.KeyS) {
			tempMove.Y += 1
		}

		playerCreature.Move(tempMove)
		playerCreature.Draw()

		rl.EndDrawing()
	}
}
