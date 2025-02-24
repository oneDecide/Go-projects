package main

import (
	"fmt"
	"math/rand"

	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	var playerX float32 = 25
	var playerY float32 = 100
	var playerSpeed float32 = 500
	var playerSize float32 = 20
	var playerPoints int = 0

	var pipeSpeed float32 = 700
	var pipeX float32 = 1250
	var pipeY float32 = 0
	var pipeSizeX float32 = 20
	var pipeSizeY float32 = 400
	var alive bool = true

	rl.InitWindow(800, 450, "2-bit Flappy Bird")
	defer rl.CloseWindow()

	rl.SetTargetFPS(200)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if alive {
			scoreText := "Score: " + strconv.Itoa(playerPoints)
			rl.DrawText(scoreText, 5, 5, 50, rl.SkyBlue)
			rl.ClearBackground(rl.Black)
			//move pipe and draw pipes
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
			if rl.IsKeyDown(rl.KeyS) && playerY < 430 {
				playerY += playerSpeed * rl.GetFrameTime()
			}

			//check intersection
			if rl.CheckCollisionRecs(playerRect, detectPipe) {
				alive = false
			}
			if rl.CheckCollisionRecs(playerRect, detect2Pipe) {
				alive = false
			}

			if pipeX < -20 {
				playerPoints++
				pipeX = 850
				fmt.Println("Y pos: ", pipeY)
				pipeY = 30 + (rand.Float32() * 420)
			}
		}
		if alive == false {
			rl.ClearBackground(rl.Red)
			rl.DrawText("GAME OVER:", 0, 10, 50, rl.Gray)
			scoreText := "Score: " + strconv.Itoa(playerPoints)
			rl.DrawText(scoreText, 0, 60, 50, rl.Gray)
			rl.DrawText("Press [ R ] to try again!", 10, 120, 40, rl.Gray)
			if rl.IsKeyDown(rl.KeyR) {
				pipeX = 1000
				alive = true
				playerPoints = 0
			}
		}

		rl.EndDrawing()
	}
}
