package main

import (
	"PhysicsEngine/ndphysics"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1920, 1080, "Snowball Madness")
	backgroundColor := rl.NewColor(47, 78, 128, 255)
	defer rl.CloseWindow()

	rl.SetTargetFPS(120)

	//PART 1, basic intersection
	//snowBall1 := ndphysics.NewProjectile(50, rl.NewVector2(0, 250), 10, rl.NewVector2(200, 0))
	//snowBall2 := ndphysics.NewProjectile(50, rl.NewVector2(1200, 200), 1000, rl.NewVector2(-20, 0))

	//PART 1 testing our existing snowballs
	simulation := ndphysics.NewSimulation()
	//simulation.AddPhysicsBody(&snowBall1.PhysicsBody)
	//simulation.AddPhysicsBody(&snowBall2.PhysicsBody)

	//Part 2add a bunch of random snowballs
	var snowBallSize float32 = 20
	var snowBallMass float32 = 10
	snowballs := make([]*ndphysics.Projectile, 0, 100)
	for i := 0; i < 100; i++ {
		randomX := rl.GetRandomValue(0, int32(rl.GetScreenWidth()))
		randomY := rl.GetRandomValue(0, int32(rl.GetScreenHeight()))
		randomPos := rl.NewVector2(float32(randomX), float32(randomY))
		randomVelX := rl.GetRandomValue(-20, 20)
		randomVelY := rl.GetRandomValue(-20, 20)
		randomVel := rl.NewVector2(float32(randomVelX), float32(randomVelY))
		snowBall := ndphysics.NewProjectile(snowBallSize, randomPos, snowBallMass, randomVel)
		simulation.AddPhysicsBody(&snowBall.PhysicsBody)
		snowballs = append(snowballs, &snowBall)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)

		//PART 1, basic intersection
		// snowBall1.PhysicsUpdate()
		// snowBall1.DrawProjectile()

		// snowBall2.PhysicsUpdate()
		// snowBall2.DrawProjectile()

		// snowBall1.CheckIntersection(&snowBall2.PhysicsBody)
		// snowBall2.CheckIntersection(&snowBall1.PhysicsBody)

		//PART 2, mass simulation
		simulation.Simualte()

		for _, snowBall := range snowballs {
			snowBall.DrawProjectile()
		}

		//snowBall1.DrawProjectile()
		//snowBall2.DrawProjectile()

		rl.EndDrawing()
	}
}
