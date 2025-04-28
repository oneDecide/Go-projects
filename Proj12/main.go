package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	gridSize     = 20
	minRoomSize  = 4
	maxRoomSize  = 10
	maxDepth     = 4
)

func main() {
	rand.Seed(time.Now().UnixNano())
	rl.InitWindow(screenWidth, screenHeight, "BSP Dungeon Generation")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	bspGen := NewBSPGenerator(screenWidth, screenHeight, gridSize)
	dungeonGrid := NewDungeonGrid(screenWidth, screenHeight, gridSize)
	bspGen.Generate()
	bspGen.CreateRooms(dungeonGrid)

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeyR) {
			bspGen.Generate()
			dungeonGrid.Reset()
			bspGen.CreateRooms(dungeonGrid)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		DrawBSPPartitions(bspGen.Root)
		dungeonGrid.Draw()
		DrawDungeonGrid(gridSize, screenWidth, screenHeight)

		rl.EndDrawing()
	}
}
