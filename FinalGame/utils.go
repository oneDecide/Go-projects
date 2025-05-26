package main

import (
	"log"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// CheckErr logs and exits on error
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// utils helper
func randomSpawn() rl.Vector2 {
	x := float32(rand.Intn(2000) - 1000)
	y := float32(rand.Intn(2000) - 1000)
	return rl.Vector2{X: x, Y: y}
}
