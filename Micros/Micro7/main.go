package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	var playerX float32 = 25
	var playerY float32 = 100
	var playerSpeed float32 = 300
	var playerSize float32 = 50

	var detectX float32 = 400
	var detectY float32 = 200
	var detectSize float32 = 100

	rl.InitWindow(800, 450, "Game!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(500)

	//playerRectangle := rl.NewRectangle(playerX, playerY, playerSize, playerSize)
	//detectRectangle := rl.NewRectangle(detectX, detectY, detectSize, detectSize)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		//draw player rectangle
		rl.DrawRectangle(int32(playerX), int32(playerY), int32(playerSize), int32(playerSize), rl.White)
		//playerX += playerSpeed * rl.GetFrameTime()
		//keyboard input
		if rl.IsKeyDown(rl.KeyW) && playerY > 0 {
			playerY -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyS) && playerY < 400 {
			playerY += playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyA) && playerX > 0 {
			playerX -= playerSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyD) && playerX < 750 {
			playerX += playerSpeed * rl.GetFrameTime()
		}

		//check intersection
		var intersecting bool = true
		if playerX > detectX+detectSize {
			intersecting = false
		} else if playerX+playerSize < detectX {
			intersecting = false
		} else if playerY > detectY+detectSize {
			intersecting = false
		} else if playerY+playerSize < detectY {
			intersecting = false
		}
		//var detectColor rl.Color = rl.White
		//if rl.CheckCollisionRecs(playerRectangle, detectRectangle) {
		//	detectColor = rl.Red
		//}
		//change color based on intersection
		//var detectColor rl.Color = rl.White
		//change color based on intersection
		var detectColor rl.Color = rl.White
		if intersecting {
			detectColor = rl.Red
		}
		//draw detect rectangle
		rl.DrawRectangle(int32(detectX), int32(detectY), int32(detectSize), int32(detectSize), detectColor)

		//draw detect rectangle
		rl.DrawRectangle(int32(detectX), int32(detectY), int32(detectSize), int32(detectSize), detectColor)
		rl.EndDrawing()
	}
}
