package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	ohmy := rl.LoadSound("audio/ohmy.ogg")
	song := rl.LoadMusicStream("audio/song.ogg")
	piano := rl.LoadMusicStream("audio/piano.wav")
	guitar := rl.LoadMusicStream("audio/guitar.wav")
	rain := rl.LoadMusicStream("audio/rain.wav")

	//if rl.GetKeyPressed()

	ohmys := make([]rl.Sound, 0, 10)
	soundIndex := 0
	for i := 0; i < 10; i++ {
		ohmys = append(ohmys, rl.LoadSoundAlias(ohmy))
	}

	// rl.SetSoundVolume(ohmy, 1)

	//rl.SetMusicVolume(rain, .75)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Use 'Q,W,E,R,T' for sound effects and music, 'Y' to stop them all", 60, 200, 20, rl.LightGray)
		rl.UpdateMusicStream(rain)
		rl.UpdateMusicStream(song)
		rl.UpdateMusicStream(piano)
		rl.UpdateMusicStream(guitar)

		if rl.IsKeyPressed(rl.KeyQ) {
			rl.PlaySound(ohmy)
			//rl.SetSoundPitch(ohmys[soundIndex], rand.Float32()+0.5)
			//rl.PlaySound(ohmys[soundIndex])
			soundIndex++
			if soundIndex >= len(ohmys) {
				soundIndex = 0
			}
		}

		if rl.IsKeyPressed(rl.KeyW) {
			rl.PlayMusicStream(rain)
		}

		if rl.IsKeyPressed(rl.KeyE) {
			rl.PlayMusicStream(song)
		}

		if rl.IsKeyPressed(rl.KeyR) {
			rl.PlayMusicStream(piano)
		}
		if rl.IsKeyPressed(rl.KeyT) {
			rl.PlayMusicStream(guitar)
		}
		if rl.IsKeyPressed(rl.KeyY) {
			rl.StopMusicStream(guitar)
			rl.StopMusicStream(piano)
			rl.StopMusicStream(song)
			rl.StopMusicStream(rain)
		}
		rl.EndDrawing()
	}

}
