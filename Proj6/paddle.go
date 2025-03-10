package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Speed float32
}

func NewPaddle() Paddle {
	return Paddle{
		Pos:   rl.NewVector2(screenWidth/2-50, screenHeight-20),
		Size:  rl.NewVector2(100, 20),
		Speed: 750,
	}
}

func (p *Paddle) Update() {
	if (rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA)) && p.Pos.X > 0 {
		p.Pos.X -= p.Speed * rl.GetFrameTime()
	}
	if (rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD)) && p.Pos.X < screenWidth-p.Size.X {
		p.Pos.X += p.Speed * rl.GetFrameTime()
	}
}

func (p *Paddle) Draw() {
	rl.DrawRectangleV(p.Pos, p.Size, rl.White)
}
