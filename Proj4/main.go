package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	var playerX float32 = 25
	var playerY float32 = 100
	var playerSpeed float32 = 300
	var playerSize float32 = 20
	var playerPoints int = 0

	var pipeSpeed float32 = 100
	var pipeX float32 = 900
	var pipeY float32 = 200
	var pipeSizeX float32 = 20
	var pipeSizeY float32 = 400

	//pipetX := pipeX
	//pipetY := pipeY + 200
	var alive bool = true

	rl.InitWindow(800, 800, "Game!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(200)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		if alive {
			//move pipe
			detectPipe := rl.NewRectangle(pipeX, pipeY, pipeSizeX, pipeSizeY)
			detect2Pipe := rl.NewRectangle(pipeX, pipeY-460, pipeSizeX, pipeSizeY)
			rl.DrawRectangleRec(detectPipe, rl.Green)
			rl.DrawRectangleRec(detect2Pipe, rl.Green)
			pipeX -= pipeSpeed * rl.GetFrameTime()
			//draw player rectangle
			playerRect := rl.NewRectangle(playerX, playerY, playerSize, playerSize)
			rl.DrawRectangleRec(playerRect, rl.Yellow)
			//keyboard input
			if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
				playerY -= playerSpeed * rl.GetFrameTime()
			}
			if rl.IsKeyDown(rl.KeyS) && playerY < 400 {
				playerY += playerSpeed * rl.GetFrameTime()
			}

			//check intersection
			if rl.CheckCollisionRecs(playerRect, detectPipe) {
				alive = false
			}
			if rl.CheckCollisionRecs(playerRect, detect2Pipe) {
				alive = false
			}

			//rl.DrawRectangle(int32(pipetX), int32(pipetY), int32(pipeSizeX), int32(pipeSizeY), rl.Green)
			if pipeX < 0 {
				playerPoints++
				pipeX = 1000
			}
		}

		rl.EndDrawing()
	}
}
