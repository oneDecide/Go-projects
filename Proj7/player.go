package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player struct {
		Position      rl.Vector2 // Center of the player's body (circle)
		Velocity      rl.Vector2
		Acceleration  float32
		Rotation      float32 // Rotation of the barrel in radians
		Speed         float32
		MaxSpeed      float32
		BodyRadius    float32 // Radius of the player's body (circle)
		BarrelLength  float32 // Length of the barrel (rectangle)
		BarrelWidth   float32 // Width of the barrel (rectangle)
		CollectedBits int     // Number of layer 1 asteroid bits collected
	}
)

func initPlayer() {
	player.Position = rl.NewVector2(400, 300)
	player.Velocity = rl.NewVector2(0, 0)
	player.Acceleration = 200 // Adjusted for delta time
	player.Rotation = 0
	player.Speed = 0
	player.MaxSpeed = 2000 // Adjusted for delta time
	player.BodyRadius = 20
	player.BarrelLength = 30
	player.BarrelWidth = 10
	player.CollectedBits = 0
}

func updatePlayer() {
	handleInput()
	updateMovement()
	updateRotation()

	// Check for collision with layer 1 asteroids
	for i := range asteroids {
		if asteroids[i].Active && asteroids[i].Layer == 1 && rl.CheckCollisionCircles(player.Position, player.BodyRadius, asteroids[i].Position, asteroids[i].Radius) {
			if player.CollectedBits < 2 {
				player.CollectedBits++
				asteroids[i].Active = false
			}
		}
	}
}

func drawPlayer() {
	// Draw the player's body (circle)
	rl.DrawCircleV(player.Position, player.BodyRadius, rl.White)

	// Calculate the barrel's position and rotation
	barrelEnd := rl.Vector2Add(player.Position, rl.Vector2Scale(rl.NewVector2(float32(math.Cos(float64(player.Rotation))), float32(math.Sin(float64(player.Rotation)))), player.BarrelLength))
	barrelCenter := rl.Vector2Lerp(player.Position, barrelEnd, 0.5)

	// Draw the barrel (rectangle)
	rl.DrawRectanglePro(
		rl.NewRectangle(barrelCenter.X, barrelCenter.Y, player.BarrelLength, player.BarrelWidth),
		rl.NewVector2(player.BarrelLength/2, player.BarrelWidth/2), // Origin (center of the rectangle)
		player.Rotation*180/math.Pi,                                // Convert radians to degrees for DrawRectanglePro
		rl.Red,
	)

	// Display collected bits
	//rl.DrawText(fmt.Sprintf("Bits: %d", player.CollectedBits), 10, 40, 20, rl.White)
}

func handleInput() {
	dt := rl.GetFrameTime() // Get delta time

	// Movement with WASD (independent of rotation)
	if rl.IsKeyDown(rl.KeyW) {
		player.Velocity.Y -= player.Acceleration * dt
	}
	if rl.IsKeyDown(rl.KeyS) {
		player.Velocity.Y += player.Acceleration * dt
	}
	if rl.IsKeyDown(rl.KeyA) {
		player.Velocity.X -= player.Acceleration * dt
	}
	if rl.IsKeyDown(rl.KeyD) {
		player.Velocity.X += player.Acceleration * dt
	}

	// Fire projectile with left mouse click
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePos := getWorldMousePosition() // Get mouse position in world coordinates
		direction := rl.Vector2Subtract(mousePos, player.Position)
		direction = rl.Vector2Normalize(direction)
		fireProjectile(player.Position, direction)
	}
}

func updateMovement() {
	dt := rl.GetFrameTime() // Get delta time

	// Limit speed
	speed := float32(math.Sqrt(float64(player.Velocity.X*player.Velocity.X + player.Velocity.Y*player.Velocity.Y)))
	if speed > player.MaxSpeed {
		scale := player.MaxSpeed / speed
		player.Velocity.X *= scale
		player.Velocity.Y *= scale
	}

	// Update position
	player.Position.X += player.Velocity.X * dt
	player.Position.Y += player.Velocity.Y * dt
}

func updateRotation() {
	// Calculate rotation to face the mouse cursor
	mousePos := getWorldMousePosition() // Get mouse position in world coordinates
	direction := rl.Vector2Subtract(mousePos, player.Position)
	player.Rotation = float32(math.Atan2(float64(direction.Y), float64(direction.X)))
}
