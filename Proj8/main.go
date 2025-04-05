package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	GridCols         = 10
	GridRows         = 10
	FarmStartX       = 3
	FarmStartY       = 3
	FarmSize         = 4
	InsideTrackWidth = 6
)

func main() {
	rl.InitWindow(800, 600, "Farm Sim")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.SetWindowState(rl.FlagWindowResizable)

	state := NewGame()
	rand.Seed(time.Now().UnixNano())

	for !rl.WindowShouldClose() {
		HandleInput(&state)
		Update(&state)
		Draw(state)
	}
}

func HandleInput(state *GameState) {
	if state.GameOver || state.GameWon {
		if rl.IsKeyPressed(rl.KeyR) {
			*state = NewGame()
		}
		return
	}

	if rl.IsKeyPressed(rl.KeyF5) { // Save
		if err := SaveGame(*state); err != nil {
			fmt.Println("Save error:", err)
		} else {
			fmt.Println("Game saved successfully")
		}
	}
	if rl.IsKeyPressed(rl.KeyF9) { // Load
		loadedState, err := LoadGame()
		if err != nil {
			fmt.Println("Load error:", err)
		} else {
			*state = loadedState
			fmt.Println("Game loaded successfully")
		}
	}

	if state.Minigame.Active {
		HandleMinigameInput(state)
		return
	}

	if state.Store.Visible {
		HandleStoreInput(state)
		return
	}

	if state.Player.IsOutside {
		HandleOutsideInput(state)
	} else {
		HandleInsideInput(state)
	}
}

func HandleOutsideInput(state *GameState) {
	prevX, prevY := state.Player.GridX, state.Player.GridY

	// Movement
	if rl.IsKeyPressed(rl.KeyW) && state.Player.GridY > 0 {
		state.Player.GridY--
	}
	if rl.IsKeyPressed(rl.KeyS) && state.Player.GridY < GridRows-1 {
		state.Player.GridY++
	}
	if rl.IsKeyPressed(rl.KeyA) && state.Player.GridX > 0 {
		state.Player.GridX--
	}
	if rl.IsKeyPressed(rl.KeyD) && state.Player.GridX < GridCols-1 {
		state.Player.GridX++
	}

	// Automatic harvesting
	if (state.Player.GridX != prevX || state.Player.GridY != prevY) &&
		state.Player.GridX >= FarmStartX && state.Player.GridX < FarmStartX+FarmSize &&
		state.Player.GridY >= FarmStartY && state.Player.GridY < FarmStartY+FarmSize {

		plotX := state.Player.GridX - FarmStartX
		plotY := state.Player.GridY - FarmStartY
		if harvestedValue := state.Crops.TryHarvest(plotX, plotY); harvestedValue > 0 {
			state.Money += harvestedValue
		}
	}

	// Watering mechanic
	if state.Player.WaterCanUses > 0 && rl.IsKeyPressed(rl.KeySpace) {
		if state.Player.GridX >= FarmStartX && state.Player.GridX < FarmStartX+FarmSize &&
			state.Player.GridY >= FarmStartY && state.Player.GridY < FarmStartY+FarmSize {

			plotX := state.Player.GridX - FarmStartX
			plotY := state.Player.GridY - FarmStartY

			if state.Crops.Plots[plotY][plotX].SeedType > 0 {
				state.Crops.Plots[plotY][plotX].Watered = true
				state.Player.WaterCanUses--
			}
		}
	}

	// House entrance
	if state.Player.GridX == 4 && state.Player.GridY == 9 && rl.IsKeyPressed(rl.KeyE) {
		state.Player.IsOutside = false
		state.Player.GridX = 0
	}

	// Water well
	if state.Player.GridX >= 7 && state.Player.GridX <= 8 &&
		state.Player.GridY >= 7 && state.Player.GridY <= 8 &&
		rl.IsKeyPressed(rl.KeyE) {
		state.Player.WaterCanUses = 5
	}

	if state.Player.WaterCanUses > 0 && rl.IsKeyPressed(rl.KeySpace) {
		if state.Player.GridX >= FarmStartX && state.Player.GridX < FarmStartX+FarmSize &&
			state.Player.GridY >= FarmStartY && state.Player.GridY < FarmStartY+FarmSize {

			plotX := state.Player.GridX - FarmStartX
			plotY := state.Player.GridY - FarmStartY

			if state.Crops.Plots[plotY][plotX].SeedType > 0 {
				state.Crops.Plots[plotY][plotX].Watered = true

				fmt.Println("Watered crop! Uses left:", state.Player.WaterCanUses)
			}
		}
	}

	// Update water well interaction
	if state.Player.GridX >= 7 && state.Player.GridX <= 8 &&
		state.Player.GridY >= 7 && state.Player.GridY <= 8 &&
		rl.IsKeyPressed(rl.KeyE) {
		state.Player.WaterCanUses = 5
		fmt.Println("Water can refilled!")
	}

	// Planting
	if state.Player.GridX >= FarmStartX && state.Player.GridX < FarmStartX+FarmSize &&
		state.Player.GridY >= FarmStartY && state.Player.GridY < FarmStartY+FarmSize {
		plotX := state.Player.GridX - FarmStartX
		plotY := state.Player.GridY - FarmStartY

		if rl.IsKeyPressed(rl.KeyOne) && state.Player.Seeds[1] > 0 {
			StartMinigame(state, plotX, plotY, 1)
			state.Player.Seeds[1]--
		} else if rl.IsKeyPressed(rl.KeyTwo) && state.Player.Seeds[2] > 0 {
			StartMinigame(state, plotX, plotY, 2)
			state.Player.Seeds[2]--
		} else if rl.IsKeyPressed(rl.KeyThree) && state.Player.Seeds[3] > 0 {
			StartMinigame(state, plotX, plotY, 3)
			state.Player.Seeds[3]--
		}
	}
}

func HandleInsideInput(state *GameState) {
	// Movement
	if rl.IsKeyPressed(rl.KeyA) && state.Player.GridX > 0 {
		state.Player.GridX--
	}
	if rl.IsKeyPressed(rl.KeyD) && state.Player.GridX < InsideTrackWidth-1 {
		state.Player.GridX++
	}

	// Bed interaction with cooldown check
	if state.Player.GridX == 5 && rl.IsKeyPressed(rl.KeyE) && !state.ShowDay {
		state.Money -= 20
		state.Day++
		state.ShowDay = true
		state.DayStart = time.Now()
		state.Crops.GrowAll()

		if state.Money <= -50 {
			state.GameOver = true
		}
	}

	// Computer
	if state.Player.GridX == 2 && rl.IsKeyPressed(rl.KeyE) {
		state.Store.Visible = true
	}

	// Exit
	if state.Player.GridX == 0 && rl.IsKeyPressed(rl.KeyE) {
		state.Player.IsOutside = true
		state.Player.GridX = 4
		state.Player.GridY = 9
	}
}

func Update(state *GameState) {
	if state.ShowDay && time.Since(state.DayStart) > 2*time.Second {
		state.ShowDay = false
	}

	if state.Money >= 500 {
		state.GameWon = true
	}
}
