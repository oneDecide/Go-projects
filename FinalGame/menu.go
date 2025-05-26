package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Menu draw/update
func UpdateMainMenu(g *Game) {
	if rl.IsKeyPressed(rl.KeyEnter) {
		g.state = StatePlaying
		g.audio.PlayMusic()
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		rl.CloseWindow()
	}
}

func DrawMainMenu() {
	rl.DrawText("Press ENTER to Start", 440, 340, 24, rl.Black)
	rl.DrawText("Press Q to Quit", 480, 380, 24, rl.Black)
}

func UpdateQuitMenu(g *Game) {
	if rl.IsKeyPressed(rl.KeyEnter) {
		g.state = StatePlaying
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		rl.CloseWindow()
	}
}

func DrawQuitMenu() {
	rl.DrawText("Resume: ENTER", 460, 340, 24, rl.Black)
	rl.DrawText("Quit: Q", 540, 380, 24, rl.Black)
}

func DrawUpgradeMenu() {
	msg := "Choose Upgrade: 1) Damage 2) Speed 3) Fire Rate"
	rl.DrawText(msg, ScreenWidth/2-200, ScreenHeight/2, 20, rl.Black)
}
