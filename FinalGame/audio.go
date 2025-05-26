package main

import rl "github.com/gen2brain/raylib-go/raylib"

// AudioManager handles background music
type AudioManager struct {
	music rl.Music
}

func NewAudioManager() *AudioManager {
	rl.InitAudioDevice()
	return &AudioManager{music: rl.LoadMusicStream("background.mp3")}
}

func (a *AudioManager) PlayMusic() { rl.PlayMusicStream(a.music) }
func (a *AudioManager) Update()    { rl.UpdateMusicStream(a.music) }
func (a *AudioManager) Close() {
	rl.UnloadMusicStream(a.music)
	rl.CloseAudioDevice()
}
