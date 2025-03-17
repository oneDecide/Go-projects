package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(1920, 1080, "Game!")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	theme := NewColorTheme(rl.NewColor(255, 255, 255, 255), rl.NewColor(255, 128, 128, 255), rl.White)

	progressBar := NewProgressBar(20, 20, 300, 100, &theme)

	slider := NewSlider(20, 150, 300, 25, 25, 50, &theme)

	pipCounter := NewPipCounter(20, 300, 20, 20, 10, 16, 32, &theme)
	pipCounter.SetPips(8)

	progressTracker := float32(0)
	direction := float32(1)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		progressBar.SetProgress(.8)
		//progressBar.DrawBar()

		slider.SetProgress(progressTracker)
		//slider.DrawSlider()

		pipCounter.DrawPipCounter()

		progressTracker += rl.GetFrameTime() * direction
		rl.EndDrawing()
	}
}
