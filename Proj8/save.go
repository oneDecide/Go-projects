package main

import (
	"encoding/gob"
	"os"
)

func SaveGame(state GameState) {
	file, err := os.Create("save.dat")
	if err != nil {
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	encoder.Encode(state)
}

func LoadGame() GameState {
	file, err := os.Open("save.dat")
	if err != nil {
		return LoadOrNewGame()
	}
	defer file.Close()

	var state GameState
	decoder := gob.NewDecoder(file)
	decoder.Decode(&state)
	return state
}

func LoadOrNewGame() GameState {
	if _, err := os.Stat("save.dat"); err == nil {
		return LoadGame()
	}
	return GameState{
		Player: Player{
			GridX:     4,
			GridY:     9,
			IsOutside: true,
			Seeds:     map[int]int{1: 3, 2: 1, 3: 0},
		},
		Crops: NewCropField(),
		Money: 100,
	}
}
