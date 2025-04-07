package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	APPLE_SIZE = 10
)

type Apple struct {
	Pos     rl.Vector2
	Carried bool
}

func (a *Apple) DrawApple() {
	rl.DrawCircle(int32(a.Pos.X), int32(a.Pos.Y), APPLE_SIZE, rl.NewColor(179, 35, 55, 255))
}

func NewApple(Pos rl.Vector2) Apple {
	return Apple{Pos, false}
}

func (a *Apple) SetPos(Pos rl.Vector2) {
	a.Pos = Pos
}
