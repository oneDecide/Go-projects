package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var alive bool = true

	rl.InitWindow(800, 450, "2-bit Moonring")
	defer rl.CloseWindow()

	rl.SetTargetFPS(200)
	MainCharacter := rl.LoadTexture("textures/StickMan.png")
	EnemyCharacter := rl.LoadTexture("textures/StickMan.png")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if alive {
			rl.ClearBackground(rl.Black)
			//check intersection
			//if rl.CheckCollisionRecs(playerRect, detectPipe) {
			//	alive = false
			//}
			DrawTextureEz(MainCharacter, rl.NewVector2(100, 400), 90, 5, rl.SkyBlue)
			DrawTextureEz(EnemyCharacter, rl.NewVector2(100, 400), 90, 5, rl.Red)
		}
		if alive == false {
			rl.ClearBackground(rl.Red)
			rl.DrawText("GAME OVER:", 0, 10, 50, rl.Gray)
			scoreText := "Score: "
			rl.DrawText(scoreText, 0, 60, 50, rl.Gray)
			rl.DrawText("Press [ R ] to try again!", 10, 120, 40, rl.Gray)
			if rl.IsKeyDown(rl.KeyR) {
				alive = true
			}
		}

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
