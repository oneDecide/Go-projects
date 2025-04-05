package main

import (
	"time"
)

type GameState struct {
	Player   Player        `json:"player"`
	Crops    CropField     `json:"crops"`
	Store    Store         `json:"store"`
	Minigame MinigameState `json:"minigame"`
	Day      int           `json:"day"`
	Money    int           `json:"money"`
	GameOver bool          `json:"gameOver"`
	GameWon  bool          `json:"gameWon"`
	ShowDay  bool          `json:"showDay"`
	DayStart time.Time     `json:"dayStart"`
}
