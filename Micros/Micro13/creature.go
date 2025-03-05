package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Pos   rl.Vector2
	Speed float32
}

func NewCreature(newPos rl.Vector2) Creature {
	newCreature := Creature{Pos: newPos, Speed: 100}
	return newCreature
}

func (c Creature) DrawCreature() {
	rl.DrawCircle(int32(c.Pos.X), int32(c.Pos.Y), 20, rl.White)
}

func (c *Creature) MoveCreature(dir rl.Vector2) {

	c.Pos = rl.Vector2Add(c.Pos, rl.Vector2Scale(dir, c.Speed*rl.GetFrameTime()))
}

func (c *Creature) MoveCreatureWithCamera(input rl.Vector2, angle float32) {
	rad := float64((angle) * (math.Pi / 180))

	upVec := rl.NewVector2(float32(math.Sin(rad)), float32(math.Cos(rad)))

	c.Pos = rl.Vector2Add(c.Pos, rl.Vector2Scale(upVec, input.Y*c.Speed*rl.GetFrameTime()))

	radRight := float64(angle+90) * (math.Pi / 180)
	rightVec := rl.NewVector2(float32(math.Sin(radRight)), float32(math.Cos(radRight)))
	// rightVec := rl.NewVector2(upVec.Y, -upVec.X)
	c.Pos = rl.Vector2Add(c.Pos, rl.Vector2Scale(rightVec, input.X*c.Speed*rl.GetFrameTime()))
}
