package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MinigameState struct {
	Active     bool           `json:"active"`
	Targets    []rl.Rectangle `json:"-"`
	StartTime  time.Time      `json:"startTime"`
	MaxTargets int            `json:"maxTargets"`
	Collected  int            `json:"collected"`
	PlotX      int            `json:"plotX"`
	PlotY      int            `json:"plotY"`
	SeedType   int            `json:"seedType"`
}

func StartMinigame(state *GameState, x, y, seedType int) {
	state.Minigame = MinigameState{
		Active:     true,
		StartTime:  time.Now(),
		MaxTargets: getMaxTargets(seedType),
		PlotX:      x,
		PlotY:      y,
		SeedType:   seedType,
	}
	generateTargets(&state.Minigame)
}

func getMaxTargets(seedType int) int {
	switch seedType {
	case 1:
		return 2
	case 2:
		return 5
	case 3:
		return 10
	}
	return 0
}

func generateTargets(m *MinigameState) {
	m.Targets = make([]rl.Rectangle, m.MaxTargets)
	for i := range m.Targets {
		m.Targets[i] = rl.NewRectangle(
			rand.Float32()*(float32(rl.GetScreenWidth())-50),
			rand.Float32()*(float32(rl.GetScreenHeight())-50),
			40,
			40,
		)
	}
}

func HandleMinigameInput(state *GameState) {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePos := rl.GetMousePosition()
		for i, target := range state.Minigame.Targets {
			if rl.CheckCollisionPointRec(mousePos, target) {
				state.Minigame.Targets = append(state.Minigame.Targets[:i], state.Minigame.Targets[i+1:]...)
				state.Minigame.Collected++
				break
			}
		}
	}

	if (time.Since(state.Minigame.StartTime) > 3*time.Second) || (state.Minigame.Collected == state.Minigame.MaxTargets) {
		bonus := state.Minigame.Collected * 5
		state.Crops.Plant(state.Minigame.PlotX, state.Minigame.PlotY, state.Minigame.SeedType, bonus)
		state.Minigame.Active = false
	}
}
