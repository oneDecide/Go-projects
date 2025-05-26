package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Diamond struct {
	Pos rl.Vector2
}

func NewDiamond(pos rl.Vector2) *Diamond { return &Diamond{Pos: pos} }

func (d *Diamond) Draw() {
	rl.DrawRectangleV(d.Pos, rl.Vector2{X: 10, Y: 10}, rl.Purple)
}

func (d *Diamond) CollidesPlayer(p *Player) bool {
	return rl.CheckCollisionRecs(
		rl.NewRectangle(d.Pos.X, d.Pos.Y, 10, 10),
		rl.NewRectangle(p.Pos.X-20, p.Pos.Y-20, 40, 40),
	)
}
