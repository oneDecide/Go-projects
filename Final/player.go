package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position     rl.Vector2
	Health       float32
	Speed        float32
	FireRate     float32
	DamageLevel  int
	Sides        int
	XP           int
	NextLevelXP  int
	Rotation     float32
	LastShotTime float32
	Color        rl.Color
}

func NewPlayer() *Player {
	return &Player{
		Position:    rl.NewVector2(0, 0),
		Health:      100,
		Speed:       3,
		FireRate:    120,
		DamageLevel: 1,
		Sides:       1,
		NextLevelXP: 100,
		Color:       rl.Blue,
	}
}

func (p *Player) Move(direction rl.Vector2) {
	if direction.X != 0 || direction.Y != 0 {
		dir := rl.Vector2Normalize(direction)
		p.Position.X += dir.X * p.Speed
		p.Position.Y += dir.Y * p.Speed
	}
}

func (p *Player) Shoot(direction rl.Vector2) []*Projectile {
	projectiles := make([]*Projectile, 0)
	p.LastShotTime = float32(rl.GetTime())

	angleStep := 2 * math.Pi / float32(p.Sides)
	for i := 0; i < p.Sides; i++ {
		angle := float32(i)*angleStep + p.Rotation
		dir := rl.NewVector2(
			float32(math.Cos(float64(angle))),
			float32(math.Sin(float64(angle))),
		)
		projectiles = append(projectiles, NewProjectile(p.Position, dir, p.DamageLevel, false))
	}
	return projectiles
}

func (p *Player) Draw() {
	const radius = 20

	// For level 1 (arrow)
	if p.Sides == 1 {
		// Draw an arrow
		tip := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)))*radius,
				float32(math.Sin(float64(p.Rotation)))*radius,
			),
		)

		leftWing := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)+2.5))*radius*0.7,
				float32(math.Sin(float64(p.Rotation)+2.5))*radius*0.7,
			),
		)

		rightWing := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)-2.5))*radius*0.7,
				float32(math.Sin(float64(p.Rotation)-2.5))*radius*0.7,
			),
		)

		points := []rl.Vector2{tip, leftWing, p.Position, rightWing}
		rl.DrawTriangleFan(points, p.Color)
		rl.DrawLineV(tip, leftWing, rl.Red)
		rl.DrawLineV(leftWing, p.Position, rl.Red)
		rl.DrawLineV(p.Position, rightWing, rl.Red)
		rl.DrawLineV(rightWing, tip, rl.Red)
		return
	}

	// For level 2 (two opposing arrows)
	if p.Sides == 2 {
		// First arrow
		tip1 := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)))*radius,
				float32(math.Sin(float64(p.Rotation)))*radius,
			),
		)

		leftWing1 := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)+2.5))*radius*0.7,
				float32(math.Sin(float64(p.Rotation)+2.5))*radius*0.7,
			),
		)

		rightWing1 := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)-2.5))*radius*0.7,
				float32(math.Sin(float64(p.Rotation)-2.5))*radius*0.7,
			),
		)

		// Second arrow (opposite direction)
		tip2 := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)+math.Pi))*radius,
				float32(math.Sin(float64(p.Rotation)+math.Pi))*radius,
			),
		)

		leftWing2 := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)+math.Pi+2.5))*radius*0.7,
				float32(math.Sin(float64(p.Rotation)+math.Pi+2.5))*radius*0.7,
			),
		)

		rightWing2 := rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(p.Rotation)+math.Pi-2.5))*radius*0.7,
				float32(math.Sin(float64(p.Rotation)+math.Pi-2.5))*radius*0.7,
			),
		)

		// Draw first arrow
		points1 := []rl.Vector2{tip1, leftWing1, p.Position, rightWing1}
		rl.DrawTriangleFan(points1, p.Color)
		rl.DrawLineV(tip1, leftWing1, rl.Red)
		rl.DrawLineV(leftWing1, p.Position, rl.Red)
		rl.DrawLineV(p.Position, rightWing1, rl.Red)
		rl.DrawLineV(rightWing1, tip1, rl.Red)

		// Draw second arrow
		points2 := []rl.Vector2{tip2, leftWing2, p.Position, rightWing2}
		rl.DrawTriangleFan(points2, p.Color)
		rl.DrawLineV(tip2, leftWing2, rl.Red)
		rl.DrawLineV(leftWing2, p.Position, rl.Red)
		rl.DrawLineV(p.Position, rightWing2, rl.Red)
		rl.DrawLineV(rightWing2, tip2, rl.Red)
		return
	}

	// For levels 3-12 (regular polygons)
	numSides := p.Sides
	if numSides > 12 {
		numSides = 12 // Cap at dodecagon (12 sides)
	}

	// Create vertices for the polygon
	vertices := make([]rl.Vector2, numSides)

	// For odd-sided polygons, one vertex points forward
	// For even-sided polygons, we offset half a vertex so an edge faces forward
	var rotationOffset float32 = 0
	if numSides%2 == 0 {
		// Even-sided polygon - point an edge forward
		rotationOffset = float32(math.Pi / float64(numSides))
	}

	// Generate vertices for the polygon
	for i := 0; i < numSides; i++ {
		angle := float32(2*math.Pi*float64(i)/float64(numSides)) + p.Rotation + rotationOffset
		vertices[i] = rl.Vector2Add(
			p.Position,
			rl.NewVector2(
				float32(math.Cos(float64(angle)))*radius,
				float32(math.Sin(float64(angle)))*radius,
			),
		)
	}

	// Draw the filled polygon
	rl.DrawTriangleFan(vertices, p.Color)

	// Draw the outline
	for i := 0; i < numSides; i++ {
		next := (i + 1) % numSides
		rl.DrawLineV(vertices[i], vertices[next], rl.Red)
	}

	// Debug: Draw a line showing the shooting direction
	frontPoint := rl.Vector2Add(
		p.Position,
		rl.NewVector2(
			float32(math.Cos(float64(p.Rotation)))*radius*1.3,
			float32(math.Sin(float64(p.Rotation)))*radius*1.3,
		),
	)
	rl.DrawLineEx(p.Position, frontPoint, 3.0, rl.Yellow)
}
