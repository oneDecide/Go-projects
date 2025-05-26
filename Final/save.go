package main

import (
	"encoding/json"
	"os"
)

type SaveState struct {
	Player            Player
	Enemies           []Enemy
	PlayerProjectiles []Projectile
	EnemyProjectiles  []Projectile
	Diamonds          []Diamond
	LevelTimer        float32
	SpawnTimer        float32
	SpawnRate         float32
	TotalGameTime     float32
}

func (g *Game) SaveGame() {
	save := SaveState{
		Player:            *g.Player,
		Enemies:           make([]Enemy, 0, len(g.Enemies)),
		PlayerProjectiles: make([]Projectile, 0, len(g.PlayerProjectiles)),
		EnemyProjectiles:  make([]Projectile, 0, len(g.EnemyProjectiles)),
		Diamonds:          make([]Diamond, 0, len(g.Diamonds)),
		LevelTimer:        g.LevelTimer,
		SpawnTimer:        g.SpawnTimer,
		SpawnRate:         g.SpawnRate,
		TotalGameTime:     g.TotalGameTime,
	}

	for _, e := range g.Enemies {
		save.Enemies = append(save.Enemies, *e)
	}

	for _, p := range g.PlayerProjectiles {
		save.PlayerProjectiles = append(save.PlayerProjectiles, *p)
	}
	for _, p := range g.EnemyProjectiles {
		save.EnemyProjectiles = append(save.EnemyProjectiles, *p)
	}

	for _, d := range g.Diamonds {
		save.Diamonds = append(save.Diamonds, *d)
	}

	data, err := json.Marshal(save)
	if err != nil {
		return
	}

	os.WriteFile("savegame.json", data, 0644)
}

func (g *Game) LoadGame() {
	data, err := os.ReadFile("savegame.json")
	if err != nil {
		return
	}

	var save SaveState
	err = json.Unmarshal(data, &save)
	if err != nil {
		return
	}

	g.Player = &save.Player

	g.Enemies = make([]*Enemy, 0, len(save.Enemies))
	for _, e := range save.Enemies {
		copy := e
		g.Enemies = append(g.Enemies, &copy)
	}

	g.PlayerProjectiles = make([]*Projectile, 0, len(save.PlayerProjectiles))
	for _, p := range save.PlayerProjectiles {
		copy := p
		g.PlayerProjectiles = append(g.PlayerProjectiles, &copy)
	}

	g.EnemyProjectiles = make([]*Projectile, 0, len(save.EnemyProjectiles))
	for _, p := range save.EnemyProjectiles {
		copy := p
		g.EnemyProjectiles = append(g.EnemyProjectiles, &copy)
	}

	g.Diamonds = make([]*Diamond, 0, len(save.Diamonds))
	for _, d := range save.Diamonds {
		copy := d
		g.Diamonds = append(g.Diamonds, &copy)
	}

	g.LevelTimer = save.LevelTimer
	g.SpawnTimer = save.SpawnTimer
	g.SpawnRate = save.SpawnRate
	g.TotalGameTime = save.TotalGameTime
	g.Camera.Target = g.Player.Position
}
