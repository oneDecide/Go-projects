package main

import (
	"math/rand/v2"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1600, 900, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	//create our animations
	idleAnimation := NewAnimation("idle", rl.LoadTexture("sprites/idle.png"), 4, .5)
	walkAnimation := NewAnimation("walk", rl.LoadTexture("sprites/walk.png"), 6, .15)

	animationFSM := NewAnimationFSM()
	animationFSM.AddAnimation(walkAnimation)
	animationFSM.AddAnimation(idleAnimation)
	animationFSM.ChangeAnimationState("idle")

	backgroundColor := rl.NewColor(35, 89, 104, 255)
	creatureColor := rl.White

	playerCreature := NewCreature(creatureColor, animationFSM)

	scoreZone := ScoreZone{rl.NewVector2(0, 0), 0}

	worldApples := make([]*Apple, 0, 10)

	for i := 0; i < 100; i++ {
		apple := NewApple(rl.NewVector2(rand.Float32()*1000, rand.Float32()*1000))
		worldApples = append(worldApples, &apple)
	}

	aiSlice := make([]*AppleThiefAI, 0, 10)
	creatureSlice := make([]*Creature, 0, 10)

	for i := 0; i < 9; i++ {
		aiCreature := NewCreature(creatureColor, animationFSM)
		aiCreature.Pos = rl.NewVector2(rand.Float32()*1000, rand.Float32()*1000)
		ai := NewAppleThiefAI(&aiCreature, &scoreZone, &worldApples)
		aiSlice = append(aiSlice, ai)
		creatureSlice = append(creatureSlice, &aiCreature)
	}

	creatureSlice = append(creatureSlice, &playerCreature)

	// Set camera offset to screen center
	camera := rl.NewCamera2D(
		rl.NewVector2(float32(rl.GetScreenWidth())/2.0, float32(rl.GetScreenHeight())/2.0), // Offset (screen center)
		playerCreature.Pos, // Initial target position
		0.0,                // Rotation
		1.0,                // Zoom
	)

	for !rl.WindowShouldClose() {
		moveDir := rl.NewVector2(0, 0)
		if rl.IsKeyDown(rl.KeyW) {
			moveDir.Y -= 1
		}
		if rl.IsKeyDown(rl.KeyS) {
			moveDir.Y += 1
		}
		if rl.IsKeyDown(rl.KeyA) {
			moveDir.X -= 1
		}
		if rl.IsKeyDown(rl.KeyD) {
			moveDir.X += 1
		}

		if rl.IsKeyPressed(rl.KeyE) {
			playerCreature.GatherApples(&worldApples)
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			playerCreature.DepositApple(&scoreZone)
		}

		playerCreature.MoveCreature(moveDir)

		camera.Target = playerCreature.Pos

		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)

		rl.BeginMode2D(camera)

		scoreZone.DrawScoreZone()

		for _, c := range creatureSlice {
			c.DrawCreature()
		}

		for _, ai := range aiSlice {
			ai.Tick()
		}

		for _, a := range worldApples {
			a.DrawApple()
		}

		rl.EndMode2D()

		rl.DrawText(strconv.Itoa(scoreZone.Points), 0, 0, 20, rl.White)
		rl.EndDrawing()
	}
}
