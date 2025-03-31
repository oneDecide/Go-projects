package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 600, "Platformer Example")
	defer rl.CloseWindow()

	meepis := Creature{
		Pos:      rl.NewVector2(400, 100),
		Vel:      rl.NewVector2(0, 0),
		Size:     rl.NewVector2(50, 50),
		FeetSize: rl.NewVector2(50, 10),
		Color:    rl.Red,
	}

	blocker := Block{
		Pos:   rl.NewVector2(325, 400),
		Size:  rl.NewVector2(100, 100),
		Color: rl.Gray,
	}

	blocker2 := Block{
		Pos:   rl.NewVector2(325, 225),
		Size:  rl.NewVector2(100, 100),
		Color: rl.Gray,
	}

	blockers := make([]Block, 0)
	blockers = append(blockers, blocker)

	gravity := rl.NewVector2(0, 980)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for _, blocker := range blockers {
			blocker.DrawBlock()
			CheckPlatformCollision(&meepis, blocker)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			meepis.Jump(blockers)
		}
		if rl.IsKeyDown(rl.KeyA) {
			meepis.ApplyGravity(rl.NewVector2(-100, 0))
		}
		if rl.IsKeyDown(rl.KeyD) {
			meepis.ApplyGravity(rl.NewVector2(100, 0))
		}
		if rl.IsKeyPressed(rl.KeyQ) {
			blockers = append(blockers, blocker2)
		}

		meepis.ApplyGravity(gravity)
		meepis.UpdateCreature()
		meepis.DrawCreature()
		rl.EndDrawing()
	}
}
