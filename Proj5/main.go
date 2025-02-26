package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var alive bool = true

	rl.InitWindow(800, 450, "2-bit Flappy Bird")
	defer rl.CloseWindow()

	rl.SetTargetFPS(200)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if alive {
			rl.ClearBackground(rl.Black)
			//check intersection
			//if rl.CheckCollisionRecs(playerRect, detectPipe) {
			//	alive = false
			//}
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
