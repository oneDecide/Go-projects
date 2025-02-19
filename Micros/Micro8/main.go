package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Game!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(500)
	MainCharacter := rl.LoadTexture("textures/StickMan.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		DrawTextureEz(MainCharacter, rl.NewVector2(100, 400), 90, 5, rl.White)
		DrawTextureEz(MainCharacter, rl.NewVector2(775, 100), 0, 3, rl.Yellow)
		DrawTextureEz(MainCharacter, rl.NewVector2(350, 200), 180, 10, rl.Blue)

		rl.EndDrawing()
	}
}

func DrawTextureEz(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)
	rl.DrawTexturePro(texture, sourceRect,
		destRect,
		origin, angle, color)
}
