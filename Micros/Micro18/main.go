package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1920, 1080, "Platformer Example")
	defer rl.CloseWindow()

	//create our animations
	idleAnimation := NewAnimation("idle", rl.LoadTexture("sprites/idle.png"), 4, .2)
	walkAnimation := NewAnimation("walk", rl.LoadTexture("sprites/walk.png"), 6, .075)
	jumpAnimation := NewAnimation("jump", rl.LoadTexture("sprites/jump.png"), 6, .075)
	jumpAnimation.Loop = false //jump doesn't loop

	animationFSM := NewAnimationFSM()
	animationFSM.AddAnimation(walkAnimation)
	animationFSM.AddAnimation(jumpAnimation)
	animationFSM.AddAnimation(idleAnimation)
	animationFSM.ChangeAnimationState("idle")

	meepis := Creature{
		Pos:          rl.NewVector2(380, 100),
		Vel:          rl.NewVector2(0, 0),
		Size:         50,
		FeetSize:     rl.NewVector2(10, 10),
		Color:        rl.Red,
		Speed:        150,
		Direction:    1,
		AnimationFSM: animationFSM,
	}

	blocker := Block{
		Pos:   rl.NewVector2(325, 200),
		Size:  rl.NewVector2(100, 100),
		Color: rl.NewColor(30, 30, 30, 255),
	}

	blockers := make([]Block, 0)
	blockers = append(blockers, blocker)

	gravity := rl.NewVector2(0, 1000)

	cam := rl.Camera2D{Offset: rl.NewVector2(0, 0), Target: rl.NewVector2(0, 0), Rotation: 0, Zoom: 2}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.BeginMode2D(cam)
		rl.ClearBackground(rl.NewColor(98, 201, 222, 255))

		for _, blocker := range blockers {
			blocker.DrawBlock()
			CheckPlatformCollision(&meepis, blocker)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			meepis.Jump(blockers)
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			blocker.Pos = rl.GetScreenToWorld2D(rl.GetMousePosition(), cam)
			blockers = append(blockers, blocker)
		}
		if rl.IsKeyDown(rl.KeyE) {
			blockers[len(blockers)-1].Pos = rl.GetScreenToWorld2D(rl.GetMousePosition(), cam)
		}

		creatureVelX := float32(0.0)
		if rl.IsKeyDown(rl.KeyD) {
			creatureVelX = 1
		}
		if rl.IsKeyDown(rl.KeyA) {
			creatureVelX = -1
		}

		meepis.Move(creatureVelX)

		meepis.ApplyGravity(gravity)
		meepis.UpdateCreature(blockers)
		meepis.DrawCreature()
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
