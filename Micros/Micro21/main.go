package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1600, 900, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	backgroundColor := rl.NewColor(171, 215, 217, 255)

	seed1 := uint64(420)
	seed2 := uint64(69)

	dungeon := NewDungeon(40, 25)
	//dungeon.Generate(seed1, seed2)

	var Camera rl.Camera2D = rl.NewCamera2D(rl.Vector2Zero(), rl.Vector2Zero(), 0, 2)

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.BeginMode2D(Camera)

		rl.ClearBackground(backgroundColor)

		if rl.IsKeyPressed(rl.KeyR) {
			seed1++
			dungeon.Generate(seed1, seed2)
		}
		dungeon.DrawDungeon()

		rl.EndMode2D()
		rl.EndDrawing()

	}
}
