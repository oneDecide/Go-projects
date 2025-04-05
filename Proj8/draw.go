package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw(state GameState) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	screenWidth := float32(rl.GetScreenWidth())
	screenHeight := float32(rl.GetScreenHeight())

	if state.GameOver || state.GameWon {
		DrawEndScreen(state, screenWidth, screenHeight)
	} else if state.ShowDay {
		DrawDayScreen(state, screenWidth, screenHeight)
	} else if state.Minigame.Active {
		DrawMinigame(state, screenWidth, screenHeight)
	} else if state.Store.Visible {
		state.Store.Draw(&state)
	} else {
		if state.Player.IsOutside {
			DrawOutside(state, screenWidth, screenHeight)
		} else {
			DrawInside(state, screenWidth, screenHeight)
		}
		DrawUI(state)
	}

	rl.EndDrawing()
}

func DrawOutside(state GameState, screenWidth, screenHeight float32) {
	cellWidth := screenWidth / GridCols
	cellHeight := screenHeight / GridRows

	// Draw grid background
	rl.DrawRectangle(0, 0, int32(screenWidth), int32(screenHeight), rl.Color{R: 240, G: 230, B: 200, A: 255})

	// Draw grid lines
	for y := 0; y <= GridRows; y++ {
		rl.DrawLineEx(
			rl.Vector2{X: 0, Y: float32(y) * cellHeight},
			rl.Vector2{X: screenWidth, Y: float32(y) * cellHeight},
			1,
			rl.ColorAlpha(rl.Gray, 0.3),
		)
	}
	for x := 0; x <= GridCols; x++ {
		rl.DrawLineEx(
			rl.Vector2{X: float32(x) * cellWidth, Y: 0},
			rl.Vector2{X: float32(x) * cellWidth, Y: screenHeight},
			1,
			rl.ColorAlpha(rl.Gray, 0.3),
		)
	}

	// Draw farmland (3x3 grid at center)
	farmRect := rl.NewRectangle(
		float32(FarmStartX)*cellWidth,
		float32(FarmStartY)*cellHeight,
		float32(FarmSize)*cellWidth,
		float32(FarmSize)*cellHeight,
	)
	rl.DrawRectangleRec(farmRect, rl.Brown)

	// Draw water well (2x2 at 7,7)
	wellRect := rl.NewRectangle(
		7*cellWidth,
		7*cellHeight,
		2*cellWidth,
		2*cellHeight,
	)
	rl.DrawRectangleRec(wellRect, rl.Blue)
	rl.DrawText("Well", int32(7*cellWidth)+5, int32(7*cellHeight)+5, 20, rl.White)

	// Draw house entrance (1x1 at 4,9)
	doorRect := rl.NewRectangle(
		4*cellWidth,
		9*cellHeight,
		cellWidth,
		cellHeight,
	)
	rl.DrawRectangleRec(doorRect, rl.DarkBrown)
	rl.DrawText("Door", int32(4*cellWidth)+5, int32(9*cellHeight)+5, 20, rl.White)

	// Draw crops
	for y := 0; y < FarmSize; y++ {
		for x := 0; x < FarmSize; x++ {
			crop := state.Crops.Plots[y][x]
			if crop.SeedType > 0 {
				drawCrop(x, y, cellWidth, cellHeight, crop)
			}
		}
	}

	// Draw player
	playerPos := rl.Vector2{
		X: (float32(state.Player.GridX) + 0.5) * cellWidth,
		Y: (float32(state.Player.GridY) + 0.5) * cellHeight,
	}
	rl.DrawCircleV(playerPos, cellWidth/3, rl.Blue)
}

func drawCrop(x, y int, cellW, cellH float32, crop Crop) {
	posX := (float32(FarmStartX+x) + 0.5) * cellW
	posY := (float32(FarmStartY+y) + 0.5) * cellH
	size := cellW * 0.4

	switch crop.SeedType {
	case 1: // Radish
		rl.DrawCircleV(rl.Vector2{X: posX, Y: posY}, size, rl.Red)
	case 2: // Wheat
		rl.DrawRectangle(
			int32(posX-size/2),
			int32(posY-size),
			int32(size),
			int32(size*2),
			rl.Gold,
		)
	case 3: // Cotton
		rl.DrawCircleV(rl.Vector2{X: posX, Y: posY}, size, rl.White)
	}
}

func DrawInside(state GameState, screenWidth, screenHeight float32) {
	cellWidth := screenWidth / InsideTrackWidth
	centerY := screenHeight / 2

	// Draw background
	rl.DrawRectangle(0, 0, int32(screenWidth), int32(screenHeight), rl.Beige)

	// Draw track
	rl.DrawRectangle(0, int32(centerY-30), int32(screenWidth), 60, rl.LightGray)

	// Draw interactables
	rl.DrawText("Exit", int32(0.2*cellWidth), int32(centerY-10), 20, rl.Black)
	rl.DrawText("Shop", int32(2.2*cellWidth), int32(centerY-10), 20, rl.Black)
	rl.DrawText("Bed", int32(5.2*cellWidth), int32(centerY-10), 20, rl.Black)

	// Draw player
	rl.DrawCircleV(
		rl.Vector2{X: (float32(state.Player.GridX) + 0.5) * cellWidth, Y: centerY},
		cellWidth/3,
		rl.Blue,
	)
}

func DrawUI(state GameState) {
	rl.DrawText(fmt.Sprintf("Money: $%d", state.Money), 10, 10, 20, rl.DarkGreen)
	rl.DrawText(fmt.Sprintf("Day: %d", state.Day), 10, 40, 20, rl.DarkBlue)
	rl.DrawText(fmt.Sprintf("Seeds: R(%d) W(%d) C(%d)",
		state.Player.Seeds[1],
		state.Player.Seeds[2],
		state.Player.Seeds[3]),
		10, 70, 20, rl.DarkBrown)
	rl.DrawText("WASD to move, E to interact", 10, 100, 20, rl.DarkGray)
	rl.DrawText(fmt.Sprintf("Water: %d", state.Player.WaterCanUses), 10, 130, 20, rl.Blue)
}

func DrawMinigame(state GameState, screenWidth, screenHeight float32) {
	// Dark background
	rl.DrawRectangle(0, 0, int32(screenWidth), int32(screenHeight), rl.Fade(rl.Black, 0.5))

	// Timer
	elapsed := time.Since(state.Minigame.StartTime)
	remaining := 3 - int(elapsed.Seconds())
	rl.DrawText(fmt.Sprintf("Time: %d", remaining),
		int32(screenWidth)/2-40,
		int32(screenHeight)/2-100,
		30,
		rl.White)

	// Targets
	for _, target := range state.Minigame.Targets {
		rl.DrawRectangleRec(target, rl.Red)
	}

	// Score
	rl.DrawText(fmt.Sprintf("Clicked: %d/%d", state.Minigame.Collected, state.Minigame.MaxTargets),
		int32(screenWidth)/2-50,
		int32(screenHeight)/2-50,
		20,
		rl.White)
}

func DrawEndScreen(state GameState, screenWidth, screenHeight float32) {
	rl.ClearBackground(rl.Black)
	text := "GAME OVER"
	color := rl.Red
	if state.GameWon {
		text = "YOU WIN!"
		color = rl.Gold
	}

	textWidth := rl.MeasureText(text, 40)
	rl.DrawText(text,
		int32(screenWidth)/2-textWidth/2,
		int32(screenHeight)/2-50,
		40,
		color)

	rl.DrawText("Press R to restart",
		int32(screenWidth)/2-rl.MeasureText("Press R to restart", 20)/2,
		int32(screenHeight)/2+50,
		20,
		rl.White)
}

func DrawDayScreen(state GameState, screenWidth, screenHeight float32) {
	rl.ClearBackground(rl.SkyBlue)
	rl.DrawText("Next Day...",
		int32(screenWidth)/2-rl.MeasureText("Next Day...", 30)/2,
		int32(screenHeight)/2-20,
		30,
		rl.White)
	rl.DrawText(fmt.Sprintf("Rent Paid: $%d", 20),
		int32(screenWidth)/2-rl.MeasureText(fmt.Sprintf("Rent Paid: $%d", 20), 20)/2,
		int32(screenHeight)/2+20,
		20,
		rl.White)
}
