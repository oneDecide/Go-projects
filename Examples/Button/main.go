package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1920, 1080, "Game!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	colorTheme := NewColorTheme(
		rl.NewColor(255, 255, 255, 255),
		rl.NewColor(128, 255, 255, 255),
		rl.NewColor(0, 0, 0, 255),
	)

	colorRandomizer := NewColorRandomizer(int32(rl.GetScreenWidth()/2), 300, 150)

	message := NewMessage("Hallo!")

	newButton := NewButton(0, 0, 300, 100, colorTheme)
	newButton.SetText("Button! :D", 20)
	newButton.CenterButtonX()
	newButton.CenterButtonY()

	newButton.AddOnClickFunc(colorRandomizer.Randomize)
	newButton.AddOnClickFunc(message.PrintMessage)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		newButton.UpdateButton()
		colorRandomizer.Draw()

		rl.EndDrawing()
	}
}
