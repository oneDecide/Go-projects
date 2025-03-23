package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Asteroid struct {
	Position rl.Vector2
	Velocity rl.Vector2
	Radius   float32
	Layer    int
	Color    rl.Color
	Active   bool
	Lifespan float32 // Lifespan in seconds
}

var (
	asteroids []Asteroid
)

func initAsteroids() {
	asteroids = make([]Asteroid, 0)
}

func spawnAsteroid() {
	// Spawn asteroid ~3000 pixels away from Earth
	angle := rand.Float32() * 2 * math.Pi
	distance := float32(3000)
	spawnPos := rl.Vector2Add(earth.Position, rl.Vector2Scale(rl.NewVector2(float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))), distance))

	// Random speed between 75 and 150
	speed := 75 + rand.Float32()*75
	direction := rl.Vector2Subtract(earth.Position, spawnPos)
	direction = rl.Vector2Normalize(direction)
	velocity := rl.Vector2Scale(direction, speed)

	// Random color (ensure it's not too dark)
	color := randomBrightColor()

	// Random layer between 4 and 2
	layer := rand.Intn(3) + 2 // 2, 3, or 4

	asteroid := Asteroid{
		Position: spawnPos,
		Velocity: velocity,
		Radius:   30, // Adjust size as needed
		Layer:    layer,
		Color:    color,
		Active:   true,
		Lifespan: 0, // No lifespan for spawned asteroids
	}
	asteroids = append(asteroids, asteroid)
}

func randomBrightColor() rl.Color {
	for {
		color := rl.Color{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}

		// Ensure the color is not too dark (brightness threshold)
		brightness := float32(color.R)*0.299 + float32(color.G)*0.587 + float32(color.B)*0.114
		if brightness > 100 { // Adjust threshold as needed
			return color
		}
	}
}

func updateAsteroids() {
	dt := rl.GetFrameTime()

	for i := range asteroids {
		if asteroids[i].Active {
			// Update position
			asteroids[i].Position.X += asteroids[i].Velocity.X * dt
			asteroids[i].Position.Y += asteroids[i].Velocity.Y * dt

			// Check for collision with Earth (before other logic)
			if asteroids[i].Layer > 1 && rl.CheckCollisionCircles(asteroids[i].Position, asteroids[i].Radius, earth.Position, earth.Radius) {
				// Reduce Earth's health by 5
				earth.Health -= 5
				asteroids[i].Active = false // Deactivate the asteroid
				continue                    // Skip further updates for this asteroid
			}

			// Update lifespan (if applicable)
			if asteroids[i].Lifespan > 0 {
				asteroids[i].Lifespan -= dt
				if asteroids[i].Lifespan <= 0 {
					asteroids[i].Active = false // Destroy asteroid when lifespan depletes
					continue
				}
			}

			// Check for collision with projectiles
			for j := range projectiles {
				if projectiles[j].Active && rl.CheckCollisionCircles(asteroids[i].Position, asteroids[i].Radius, projectiles[j].Position, 5) {
					// Split asteroid
					splitAsteroid(&asteroids[i])
					projectiles[j].Active = false
					break
				}
			}
		}
	}
}

func splitAsteroid(asteroid *Asteroid) {
	if asteroid.Layer > 1 { // Only split if layer is greater than 1
		// Split into 2 smaller asteroids
		for i := 0; i < 2; i++ {
			angle := rand.Float32() * 2 * math.Pi
			direction := rl.NewVector2(float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle))))
			velocity := rl.Vector2Scale(direction, rl.Vector2Length(asteroid.Velocity)+75)

			// Determine new layer and color
			newLayer := asteroid.Layer - 1
			color := asteroid.Color
			if newLayer == 1 {
				color = rl.Red // Layer 1 asteroids are bright red
			}

			// Set fixed size for layer 1 asteroids
			radius := asteroid.Radius / 1.5
			if newLayer == 1 {
				radius = 15 // Fixed size for layer 1 asteroids (30x30 pixels)
			}

			newAsteroid := Asteroid{
				Position: asteroid.Position,
				Velocity: velocity,
				Radius:   radius,
				Layer:    newLayer,
				Color:    color,
				Active:   true,
				Lifespan: 20, // 20-second lifespan for split asteroids
			}
			asteroids = append(asteroids, newAsteroid)
		}
	}
	asteroid.Active = false
}

func drawAsteroids() {
	for _, asteroid := range asteroids {
		if asteroid.Active {
			if asteroid.Layer == 1 {
				// Draw layer 1 asteroids as squares (30x30 pixels)
				rl.DrawRectanglePro(
					rl.NewRectangle(asteroid.Position.X, asteroid.Position.Y, 20, 20),
					rl.NewVector2(15, 15), // Origin (center of the square)
					0,                     // No rotation
					asteroid.Color,
				)
			} else {
				// Draw layers 2+ as circles
				rl.DrawCircleV(asteroid.Position, asteroid.Radius, asteroid.Color)
			}
		}
	}
}
