package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1980
	screenHeight = 1080
)

var (
	score    int
	gameOver bool
	camera   rl.Camera2D
)

func initGame() {
	score = 0
	initPlayer()
	initMines()
	camera = rl.NewCamera2D(
		rl.NewVector2(screenWidth/2, screenHeight/2),
		rl.NewVector2(screenWidth/2, screenHeight/2),
		0,
		1,
	)
	gameOver = false
}

func main() {
	rl.InitWindow(screenWidth, screenHeight, "DEVIOUS MINESWEEPER")
	defer rl.CloseWindow()
	rl.SetTargetFPS(120)
	fmt.Println("yo")
	initGame()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if !gameOver {
			updateMovement()
			rl.ClearBackground(rl.Gray)

			drawPlayer()
			drawMine()
			handleInput()
			updateMine()

		} else {
			rl.DrawText("Game Over! Press R to Restart", screenWidth/2-150, screenHeight/2-20, 20, rl.Yellow)
			rl.ClearBackground(rl.DarkBrown)
			if rl.IsKeyPressed(rl.KeyR) {
				initGame()
			}
		}

		rl.DrawText(fmt.Sprintf("Score: %d", score), 20, 20, 50, rl.White)

		rl.EndDrawing()
	}
}

func getWorldMousePosition() rl.Vector2 {
	return rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
}

func GameOver() {
	gameOver = true
}
