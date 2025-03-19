package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Game!")
	defer rl.CloseWindow()

	cySprite := rl.LoadTexture("sprites/cyclops.png")
	sSprite := rl.LoadTexture("sprites/starliner.png")

	fmt.Println(cySprite)

	testCreature := NewCreature("Meepis", 100, 100, cySprite, rl.White)

	testCreature.AddItem(Item{"Lunchbox", 10})
	testCreature.AddItem(Item{"Sword", 10})

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		testCreature.DrawCreature()

		if rl.IsKeyPressed(rl.KeyR) {
			testCreature.RandomizeColor()
		}
		if rl.IsKeyPressed(rl.KeyS) {
			testCreature.Save()
		}
		if rl.IsKeyPressed(rl.KeyL) {
			testCreature.Load()
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			testCreature.Texture = sSprite
		}

		rl.EndDrawing()
	}
}
