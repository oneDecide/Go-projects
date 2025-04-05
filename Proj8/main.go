package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	GridCols         = 10
	GridRows         = 10
	FarmStartX       = 3
	FarmStartY       = 3
	FarmSize         = 3
	InsideTrackWidth = 6
)

type GameState struct {
	Player   Player
	Crops    CropField
	Store    Store
	Minigame MinigameState
	Day      int
	Money    int
	GameOver bool
	GameWon  bool
	ShowDay  bool
	DayStart time.Time
}

func main() {
	rl.InitWindow(800, 600, "Farm Sim")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	rl.SetWindowState(rl.FlagWindowResizable)

	state := LoadOrNewGame()
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
			*state = LoadOrNewGame()
		}
		return
	}

	if rl.IsKeyPressed(rl.KeyO) {
		SaveGame(*state)
	}
	if rl.IsKeyPressed(rl.KeyP) {
		*state = LoadGame()
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
func HandleStoreInput(state *GameState) {
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyX) {
		state.Store.Visible = false
	}

	// Handle mouse clicks for purchases
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		mousePos := rl.GetMousePosition()
		width := float32(rl.GetScreenWidth())
		height := float32(rl.GetScreenHeight())
		shopWidth := width * 0.6
		shopHeight := height * 0.6
		x := (width - shopWidth) / 2
		y := (height - shopHeight) / 2

		items := []struct {
			price int
			seed  int
			rect  rl.Rectangle
		}{
			{5, 1, rl.NewRectangle(x+20, y+50, shopWidth-40, 30)},
			{30, 2, rl.NewRectangle(x+20, y+90, shopWidth-40, 30)},
			{60, 3, rl.NewRectangle(x+20, y+130, shopWidth-40, 30)},
		}

		for _, item := range items {
			if rl.CheckCollisionPointRec(mousePos, item.rect) && state.Money >= item.price {
				state.Money -= item.price
				state.Player.Seeds[item.seed]++
			}
		}
	}
}
func HandleOutsideInput(state *GameState) {
	//prevX, prevY := state.Player.GridX, state.Player.GridY

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
	if rl.IsKeyPressed(rl.KeyA) && state.Player.GridX > 0 {
		state.Player.GridX--
	}
	if rl.IsKeyPressed(rl.KeyD) && state.Player.GridX < InsideTrackWidth-1 {
		state.Player.GridX++
	}

	// Bed
	if state.Player.GridX == 5 && rl.IsKeyPressed(rl.KeyE) {
		state.Money -= 20
		state.Day++
		state.ShowDay = true
		state.DayStart = time.Now()
		state.Crops.GrowAll()
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

	if state.Money <= -50 {
		state.GameOver = true
	}
	if state.Money >= 500 {
		state.GameWon = true
	}
}
