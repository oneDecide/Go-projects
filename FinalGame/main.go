package main

import rl "github.com/gen2brain/raylib-go/raylib"

var gameRef *Game

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Polygon Survivor")
	rl.SetTargetFPS(TargetFPS)

	gameRef = NewGame()
	defer gameRef.Close()

	for !rl.WindowShouldClose() {
		gameRef.Update()
		gameRef.Draw()
	}

	rl.CloseWindow()
}
