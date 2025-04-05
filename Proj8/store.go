package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Store struct {
	Visible bool
}

func (s *Store) Draw(state *GameState) {
	width := float32(rl.GetScreenWidth())
	height := float32(rl.GetScreenHeight())

	// Dark background
	rl.DrawRectangle(0, 0, int32(width), int32(height), rl.Fade(rl.Black, 0.5))

	// Shop window - using regular rectangle
	shopWidth := width * 0.6
	shopHeight := height * 0.6
	x := int32((width - shopWidth) / 2)
	y := int32((height - shopHeight) / 2)

	rl.DrawRectangle(x, y, int32(shopWidth), int32(shopHeight), rl.LightGray)
	rl.DrawRectangleLines(x, y, int32(shopWidth), int32(shopHeight), rl.DarkGray)

	// Close button
	closeRect := rl.Rectangle{X: float32(x) + shopWidth - 40, Y: float32(y) + 10, Width: 30, Height: 30}
	closeColor := rl.Maroon
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), closeRect) {
		closeColor = rl.Red
	}
	rl.DrawRectangleRec(closeRect, closeColor)
	rl.DrawText("X", int32(closeRect.X)+10, int32(closeRect.Y)+5, 20, rl.White)

	// Seed items
	items := []struct {
		name  string
		price int
		seed  int
	}{
		{"Radish Seeds ($5)", 5, 1},
		{"Wheat Seeds ($30)", 30, 2},
		{"Cotton Seeds ($60)", 60, 3},
	}

	itemY := float32(y) + 50
	for _, item := range items {
		itemRect := rl.NewRectangle(float32(x)+20, itemY, shopWidth-40, 30)
		itemColor := rl.DarkBrown
		if rl.CheckCollisionPointRec(rl.GetMousePosition(), itemRect) {
			rl.DrawRectangleRec(itemRect, rl.ColorAlpha(rl.SkyBlue, 0.3))
		}
		rl.DrawText(item.name, int32(itemRect.X)+10, int32(itemRect.Y)+5, 20, itemColor)
		itemY += 40
	}
}
