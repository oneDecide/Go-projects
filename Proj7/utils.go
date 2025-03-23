package main

import rl "github.com/gen2brain/raylib-go/raylib"

func distance(a, b rl.Vector2) float32 {
	return float32(rl.Vector2Distance(a, b))
}
