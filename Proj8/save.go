package main

import (
	"encoding/json"
	"os"
)

const saveFile = "farm_sim_save.json"

func NewGame() GameState {
	return GameState{
		Player: Player{
			GridX:        4,
			GridY:        9,
			IsOutside:    true,
			Seeds:        map[int]int{1: 3, 2: 1, 3: 0},
			WaterCanUses: 0,
		},
		Crops:    NewCropField(),
		Money:    100,
		Day:      1,
		GameOver: false,
		GameWon:  false,
	}
}

func SaveGame(state GameState) error {
	file, err := os.Create(saveFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(state)
}

func LoadGame() (GameState, error) {
	file, err := os.Open(saveFile)
	if err != nil {
		return GameState{}, err
	}
	defer file.Close()

	var state GameState
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&state)
	return state, err
}
