package main

import (
	"PhysicsEngine/ndphysics"
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1280, 720, "Snowball")

	backgroundColor := rl.NewColor(47, 78, 128, 255)

	numSnowballs := 100

	defer rl.CloseWindow()

	rl.SetTargetFPS(144)

	snowballs := make([]ndphysics.Projectile, numSnowballs)

	for i := 0; i < numSnowballs; i++ {
		snowballs[i] = ndphysics.NewProjectile(20, rl.NewVector2(float32(rand.IntN(1280)), float32(rand.IntN(720))), rl.NewVector2(float32(rand.IntN(1000)-500), float32(rand.IntN(1000)-500)))
		snowballs[i].Gravity = rl.NewVector2(0, float32(rand.IntN(1000)-500))
	}

	//snowBall1 := ndphysics.NewProjectile(20, rl.NewVector2(200, 200), rl.NewVector2(200, -500))
	//snowBall1.Gravity = rl.NewVector2(0, 500)

	//snowBall2 := ndphysics.NewProjectile(20, rl.NewVector2(400, 200), rl.NewVector2(-200, -500))
	//snowBall2.Gravity = rl.NewVector2(0, 250)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)

		for i := 0; i < numSnowballs; i++ {
			snowballs[i].PhysicsUpdate()
			snowballs[i].DrawProjectile()
		}

		rl.EndDrawing()
	}
}
