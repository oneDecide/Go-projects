package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ColorRandomizer struct {
	X            int32
	Y            int32
	radius       float32
	currentColor rl.Color
}
func (cr *ColorRandomizer) Randomize() {
	cr.currentColor = rl.NewColor(uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255)
	fmt.Println(cr.currentColor)
}

func (cr ColorRandomizer) Draw() {
	rl.DrawCircle(cr.X, cr.Y, cr.radius, cr.currentColor)
}


func NewColorRandomizer(newX, newY int32, newRad float32) ColorRandomizer {
	c := ColorRandomizer{X: newX, Y: newY, radius: newRad, currentColor: rl.White}
	return c
}

