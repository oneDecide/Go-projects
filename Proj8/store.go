package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Store struct {
	Visible      bool
	SelectedItem int
}

func (s *Store) Draw(state *GameState) {
	if !s.Visible {
		return
	}

	width := float32(rl.GetScreenWidth())
	height := float32(rl.GetScreenHeight())

	rl.DrawRectangle(0, 0, int32(width), int32(height), rl.Fade(rl.Black, 0.5))

	shopWidth := width * 0.6
	shopHeight := height * 0.5
	x := (width - shopWidth) / 2
	y := (height - shopHeight) / 2
	rl.DrawRectangleRec(rl.NewRectangle(x, y, shopWidth, shopHeight), rl.LightGray)
	rl.DrawRectangleLinesEx(rl.NewRectangle(x, y, shopWidth, shopHeight), 2, rl.DarkGray)

	title := "SHOP - Press X to exit"
	titleWidth := rl.MeasureText(title, 20)
	rl.DrawText(title, int32(x+shopWidth/2-float32(titleWidth)/2), int32(y+10), 20, rl.DarkBlue)

	items := []struct {
		name  string
		price int
		seed  int
	}{
		{"1. Radish Seeds - $5", 5, 1},
		{"2. Wheat Seeds - $30", 30, 2},
		{"3. Cotton Seeds - $60", 60, 3},
	}

	itemY := y + 50
	for _, item := range items {
		color := rl.DarkBrown
		if s.SelectedItem == item.seed {
			rl.DrawRectangleRec(rl.NewRectangle(x+10, itemY-5, shopWidth-20, 30), rl.ColorAlpha(rl.SkyBlue, 0.3))
			color = rl.DarkBlue
		}

		rl.DrawText(item.name, int32(x+20), int32(itemY), 20, color)
		itemY += 40
	}

	instructions := "Press 1-3 to buy seeds | X to exit"
	instructionsWidth := rl.MeasureText(instructions, 20)
	rl.DrawText(instructions, int32(x+shopWidth/2-float32(instructionsWidth)/2), int32(y+shopHeight-40), 20, rl.DarkGray)
}

func HandleStoreInput(state *GameState) {
	if rl.IsKeyPressed(rl.KeyX) {
		state.Store.Visible = false
		state.Store.SelectedItem = 0
		return
	}

	if rl.IsKeyPressed(rl.KeyOne) || rl.IsKeyPressed(rl.KeyTwo) || rl.IsKeyPressed(rl.KeyThree) {
		seedType := 0
		switch {
		case rl.IsKeyPressed(rl.KeyOne):
			seedType = 1
		case rl.IsKeyPressed(rl.KeyTwo):
			seedType = 2
		case rl.IsKeyPressed(rl.KeyThree):
			seedType = 3
		}

		state.Store.SelectedItem = seedType
		buySeed(state, seedType)
	}
}

func buySeed(state *GameState, seedType int) {
	prices := map[int]int{1: 5, 2: 30, 3: 60}
	price := prices[seedType]

	if state.Money >= price {
		state.Money -= price
		state.Player.Seeds[seedType]++
	}
}
