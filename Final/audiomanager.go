package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type AudioManager struct {
	music      rl.Music
	shotSound  rl.Sound
	collect    rl.Sound
	audioReady bool
}

func NewAudioManager() *AudioManager {
	am := &AudioManager{}
	if rl.IsAudioDeviceReady() {
		am.music = rl.LoadMusicStream("Audio/music.mp3")
		am.shotSound = rl.LoadSound("Audio/shot.wav")
		am.collect = rl.LoadSound("Audio/collect.wav")
		am.audioReady = true

		rl.PlayMusicStream(am.music)
		rl.SetMasterVolume(.12)
	}
	return am
}

func (am *AudioManager) PlayShot() {
	if am.audioReady {
		rl.PlaySound(am.shotSound)
	}
}

func (am *AudioManager) PlayCollect() {
	if am.audioReady {
		rl.PlaySound(am.collect)
	}
}

func (am *AudioManager) Update() {
	if am.audioReady {
		rl.UpdateMusicStream(am.music)
	}
}

func (am *AudioManager) Cleanup() {
	if am.audioReady {
		rl.UnloadMusicStream(am.music)
		rl.UnloadSound(am.shotSound)
		rl.UnloadSound(am.collect)
	}
}
