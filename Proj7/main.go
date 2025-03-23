package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800
	screenHeight = 600
	spawnRadius  = 3000 // Distance from Earth where asteroids spawn
)

var (
	camera        rl.Camera2D
	gameOver      bool
	spawnTimer    float32       // Timer for asteroid spawning
	spawnInterval float32 = 1.0 // Spawn an asteroid every 1 second
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Space Defender")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Initialize game objects
	initGame()

	for !rl.WindowShouldClose() {
		update()
		draw()
	}
}

func initGame() {
	// Reset game state
	gameOver = false
	initPlayer()
	initEarth(screenWidth/2, screenHeight/2)
	initProjectiles()
	initAsteroids()

	// Initialize camera
	camera = rl.NewCamera2D(
		rl.NewVector2(screenWidth/2, screenHeight/2), // Target (center of screen)
		player.Position, // Offset (player position)
		0,               // Rotation
		1,               // Zoom
	)

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
}

func update() {
	if !gameOver {
		dt := rl.GetFrameTime() // Get delta time

		updatePlayer()
		updateProjectiles()
		updateAsteroids()
		updateEarth()
		updateCamera()

		// Update spawn timer
		spawnTimer += dt

		// Spawn an asteroid every second
		if spawnTimer >= spawnInterval {
			spawnAsteroid()
			spawnTimer = 0 // Reset the timer
		}

		// Check for game over condition
		if earth.Health <= 0 {
			gameOver = true
		}
	} else {
		// Restart the game if R is pressed
		if rl.IsKeyPressed(rl.KeyR) {
			initGame()
		}
	}
}

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.BeginMode2D(camera)

	drawEarth()
	drawPlayer()
	drawProjectiles()
	drawAsteroids()

	rl.EndMode2D()

	// Draw UI (health and bits)
	drawUI()

	// Display game over message if the game is over
	if gameOver {
		rl.DrawText("Game Over! Press R to Restart", screenWidth/2-150, screenHeight/2-20, 20, rl.Red)
	}

	rl.EndDrawing()
}

func updateCamera() {
	// Camera follows player
	camera.Target = player.Position
}

// Helper function to get mouse position in world coordinates
func getWorldMousePosition() rl.Vector2 {
	return rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
}
