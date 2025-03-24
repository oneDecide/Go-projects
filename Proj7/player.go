package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	player struct {
		Position      rl.Vector2
		Velocity      rl.Vector2
		Acceleration  float32
		Rotation      float32
		Speed         float32
		MaxSpeed      float32
		BodyRadius    float32
		BarrelLength  float32
		BarrelWidth   float32
		CollectedBits int
	}
)

func initPlayer() {
	player.Position = rl.NewVector2(400, 300)
	player.Velocity = rl.NewVector2(0, 0)
	player.Acceleration = 200
	player.Rotation = 0
	player.Speed = 0
	player.MaxSpeed = 2000
	player.BodyRadius = 20
	player.BarrelLength = 30
	player.BarrelWidth = 10
	player.CollectedBits = 0
}

func updatePlayer() {
	handleInput()
	updateMovement()
	updateRotation()

	for i := range asteroids {
		if asteroids[i].Active && asteroids[i].Layer == 1 && rl.CheckCollisionCircles(player.Position, player.BodyRadius, asteroids[i].Position, asteroids[i].Radius) {
			if player.CollectedBits < 2 {
				player.CollectedBits++
				asteroids[i].Active = false
				playSoundWithPitch(pickupSound)
			}
		}
	}
}

func drawPlayer() {
	//player
	rl.DrawCircleV(player.Position, player.BodyRadius, rl.White)

	//barrel rotation
	barrelEnd := rl.Vector2Add(player.Position, rl.Vector2Scale(rl.NewVector2(float32(math.Cos(float64(player.Rotation))), float32(math.Sin(float64(player.Rotation)))), player.BarrelLength))
	barrelCenter := rl.Vector2Lerp(player.Position, barrelEnd, 0.5)

	//barrel
	rl.DrawRectanglePro(
		rl.NewRectangle(barrelCenter.X, barrelCenter.Y, player.BarrelLength, player.BarrelWidth),
		rl.NewVector2(player.BarrelLength/2, player.BarrelWidth/2),
		player.Rotation*180/math.Pi,
		rl.Red,
	)
}

func handleInput() {
	dt := rl.GetFrameTime()

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

	//lft click to fire projectile
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsKeyPressed(rl.KeySpace) {
		mousePos := getWorldMousePosition()
		direction := rl.Vector2Subtract(mousePos, player.Position)
		direction = rl.Vector2Normalize(direction)
		fireProjectile(player.Position, direction)
		playSoundWithPitch(shootSound)
	}
}

func updateMovement() {
	dt := rl.GetFrameTime()
	speed := float32(math.Sqrt(float64(player.Velocity.X*player.Velocity.X + player.Velocity.Y*player.Velocity.Y)))
	if speed > player.MaxSpeed {
		scale := player.MaxSpeed / speed
		player.Velocity.X *= scale
		player.Velocity.Y *= scale
	}

	player.Position.X += player.Velocity.X * dt
	player.Position.Y += player.Velocity.Y * dt
}

func updateRotation() {
	//rotation for ship to look at mouse
	mousePos := getWorldMousePosition()
	direction := rl.Vector2Subtract(mousePos, player.Position)
	player.Rotation = float32(math.Atan2(float64(direction.Y), float64(direction.X)))
}

func playSoundWithPitch(sound rl.Sound) {
	pitch := 1.0 + (rand.Float64()-0.5)*0.1 // Random pitch variation ±0.05
	rl.SetSoundPitch(sound, float32(pitch))
	rl.SetSoundVolume(sound, .3)
	rl.PlaySound(sound)
}
