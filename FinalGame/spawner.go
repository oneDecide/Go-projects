package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Spawner struct {
	Timer        float32
	BaseRate     float32
	SpawnRate    float32
	LastIncrease float64
	Player       *Player // store player reference
}

func NewSpawner(player *Player) *Spawner {
	return &Spawner{
		Timer:        0,
		BaseRate:     7.0,
		SpawnRate:    7.0,
		LastIncrease: rl.GetTime(),
		Player:       player,
	}
}

func (s *Spawner) Update() {
	now := rl.GetTime()
	elapsed := float32(now - s.LastIncrease)

	// Decrease spawn rate every 60s, minimum 1.5s
	if elapsed >= 60 {
		s.SpawnRate -= 1.0
		if s.SpawnRate < 1.5 {
			s.SpawnRate = 1.5
		}
		s.LastIncrease = now
	}

	s.Timer += rl.GetFrameTime()
	if s.Timer >= s.SpawnRate {
		s.Timer = 0
		enemyType := rand.Intn(2)
		if enemyType == 0 {
			gameRef.enemies = append(gameRef.enemies, NewCircleEnemy(s.Player))
		} else {
			gameRef.enemies = append(gameRef.enemies, NewArrowEnemy(s.Player))
		}
	}
}
