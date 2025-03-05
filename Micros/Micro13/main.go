package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

func main() {
	rl.InitWindow(1920, 1080, "raylib [core] example - basic window")

	playerCreature := NewCreature(rl.NewVector2(0, 0))
	playerCreature2 := NewCreature(rl.NewVector2(100, 100))

	camAngle := float32(0.0)
	camZoom := float32(1)
	camRotationSpeed := float32(50)
	camZoomSpeed := float32(5)
	char1 := true

	cam := rl.NewCamera2D(
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0,
		1,
	)
	cam.Offset = rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2))

	trees := make([]Tree, 0, 5)
	for i := 0; i < 100; i++ {
		trees = append(trees, NewTree(rl.NewVector2(float32(rand.Intn(1000)-500), float32(rand.Intn(1000)-500))))
	}

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		dir := rl.NewVector2(0, 0)
		if rl.IsKeyDown(rl.KeyW) {
			dir.Y -= 1
		}
		if rl.IsKeyDown(rl.KeyS) {
			dir.Y += 1
		}
		if rl.IsKeyDown(rl.KeyA) {
			dir.X -= 1
		}
		if rl.IsKeyDown(rl.KeyD) {
			dir.X += 1
		}

		if rl.IsKeyDown(rl.KeyQ) {
			camAngle += rl.GetFrameTime() * camRotationSpeed
		}
		if rl.IsKeyDown(rl.KeyE) {
			camAngle -= rl.GetFrameTime() * camRotationSpeed
		}
		if rl.IsKeyDown(rl.KeyZ) {
			camZoom += rl.GetFrameTime() * camZoomSpeed
		}
		if rl.IsKeyDown(rl.KeyX) {
			camZoom -= rl.GetFrameTime() * camZoomSpeed
			if camZoom < 0.1 { //zoom of 0 causes issues
				camZoom = 0.1
			}
		}
		if rl.IsKeyPressed(rl.KeyT) {
			char1 = !char1
		}
		if char1 {
			//playerCreature.MoveCreature(dir)

			playerCreature.MoveCreatureWithCamera(dir, camAngle)
			cam.Target = rl.NewVector2(float32(int32(playerCreature.Pos.X)), float32(int32(playerCreature.Pos.Y)))
			cam.Rotation = camAngle
			cam.Zoom = camZoom
		}
		if !char1 {
			//playerCreature2.MoveCreature(dir)

			playerCreature2.MoveCreatureWithCamera(dir, camAngle)
			cam.Target = rl.NewVector2(float32(int32(playerCreature2.Pos.X)), float32(int32(playerCreature2.Pos.Y)))
			cam.Rotation = camAngle
			cam.Zoom = camZoom
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(30, 90, 100, 255))
		rl.BeginMode2D(cam)

		for _, tree := range trees {
			tree.DrawTree()
		}

		playerCreature.DrawCreature()
		playerCreature2.DrawCreature()

		rl.EndMode2D()

		rl.DrawText("Tree Game", 0, 0, 20, rl.White)

		rl.EndDrawing()
	}
}
