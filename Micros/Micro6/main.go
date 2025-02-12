package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1280, 720, "raylib [core] example - basic window")

	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawRectangle(0, 0, 100, 100, rl.Red)
		rl.DrawRectangle(1180, 0, 100, 100, rl.Blue)
		rl.DrawRectangle(0, 620, 100, 100, rl.Green)
		rl.DrawRectangle(1180, 620, 100, 100, rl.Yellow)
		rl.DrawRectangle(1180/2, 620/2, 100, 100, rl.Color{155, 155, 255, 100})

		rl.EndDrawing()
	}
}
