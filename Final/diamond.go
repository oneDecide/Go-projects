package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Diamond struct {
	Position rl.Vector2
	IsActive bool
}

func NewDiamond(pos rl.Vector2) *Diamond {
	return &Diamond{
		Position: pos,
		IsActive: true,
	}
}

func (d *Diamond) Draw() {
	rl.DrawRectangleV(
		rl.Vector2Subtract(d.Position, rl.NewVector2(5, 5)),
		rl.NewVector2(10, 10),
		rl.SkyBlue,
	)
}
