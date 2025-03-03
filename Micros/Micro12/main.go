package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	ohmy := rl.LoadSound("audio/ohmy.ogg")
	rain := rl.LoadMusicStream("audio/song.ogg")

	//if rl.GetKeyPressed()

	ohmys := make([]rl.Sound, 0, 10)
	soundIndex := 0
	for i := 0; i < 10; i++ {
		ohmys = append(ohmys, rl.LoadSoundAlias(ohmy))
	}

	// rl.SetSoundVolume(ohmy, 1)
	rl.SetSoundPitch(ohmy, .75)

	//rl.SetMusicVolume(rain, .75)
	//rl.SetMusicPitch(rain, 1.25)

	rl.PlayMusicStream(rain)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.UpdateMusicStream(rain)

		if rl.IsKeyPressed(rl.KeyO) {
			//rl.PlaySound(ohmy)
			rl.SetSoundPitch(ohmys[soundIndex], rand.Float32()+0.5)
			rl.PlaySound(ohmys[soundIndex])
			soundIndex++
			if soundIndex >= len(ohmys) {
				soundIndex = 0
			}
		}

		if rl.IsKeyPressed(rl.KeyR) {
			rl.StopMusicStream(rain)
			rl.PlayMusicStream(rain)
		}

		if rl.IsKeyPressed(rl.KeyP) {
			rl.PauseMusicStream(rain)
		}
		if rl.IsKeyPressed(rl.KeyL) {
			rl.ResumeMusicStream(rain)
		}

		rl.EndDrawing()
	}

}
