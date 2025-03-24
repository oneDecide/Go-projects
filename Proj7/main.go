package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1980
	screenHeight = 1080
	spawnRadius  = 3000
)

var (
	camera           rl.Camera2D
	gameOver         bool
	spawnTimer       float32
	spawnInterval    float32 = 1.0
	music            rl.Music
	shootSound       rl.Sound
	pickupSound      rl.Sound
	hurtSound        rl.Sound
	destructionSound rl.Sound
	depositSound     rl.Sound
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Space Defender")
	defer rl.CloseWindow()
	rl.InitAudioDevice()

	music = rl.LoadMusicStream("audio/music.wav")
	shootSound = rl.LoadSound("audio/shoot.wav")
	pickupSound = rl.LoadSound("audio/pickup.wav")
	hurtSound = rl.LoadSound("audio/hurt.wav")
	destructionSound = rl.LoadSound("audio/destruction.wav")
	depositSound = rl.LoadSound("audio/deposit.wav")

	rl.PlayMusicStream(music)
	rl.SetMusicVolume(music, 0.5)

	rl.SetTargetFPS(60)

	initGame()

	//start := false
	for !rl.WindowShouldClose() {

		update()
		draw()
		rl.UpdateMusicStream(music)

		//rl.UnloadSound(shootSound)
	}
}

func initGame() {
	gameOver = false
	initPlayer()
	initEarth(screenWidth/2, screenHeight/2)
	initProjectiles()
	initAsteroids()

	camera = rl.NewCamera2D(
		rl.NewVector2(screenWidth/2, screenHeight/2),
		player.Position,
		0,
		1,
	)

	rand.Seed(time.Now().UnixNano())
}

func update() {

	if !gameOver {
		dt := rl.GetFrameTime()

		updatePlayer()
		updateProjectiles()
		updateAsteroids()
		updateEarth()
		updateCamera()

		spawnTimer += dt

		if spawnTimer >= spawnInterval {
			spawnAsteroid()
			spawnTimer = 0
		}

		if earth.Health <= 0 {
			gameOver = true
		}
	} else {
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

	drawUI()
	if gameOver {
		rl.DrawText("Game Over! Press R to Restart", screenWidth/2-150, screenHeight/2-20, 20, rl.Red)
	}

	rl.EndDrawing()
}

func updateCamera() {
	camera.Target = player.Position
}

func getWorldMousePosition() rl.Vector2 {
	return rl.GetScreenToWorld2D(rl.GetMousePosition(), camera)
}
