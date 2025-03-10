package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Ball struct {
	Pos    rl.Vector2
	Vel    rl.Vector2
	Radius float32
	Speed  float32
	Locked bool
}

func NewBall() Ball {
	return Ball{
		Pos:    rl.NewVector2(screenWidth/2, screenHeight-40),
		Vel:    rl.NewVector2(0, 0),
		Radius: 10,
		Speed:  400,
		Locked: true,
	}
}

func (b *Ball) Update(paddle *Paddle, blocks *[]Block) {

	if b.Locked {
		b.Pos.X = paddle.Pos.X + paddle.Size.X/2 + 10 - b.Radius
		b.Pos.Y = paddle.Pos.Y - b.Radius*2

		if rl.IsKeyPressed(rl.KeySpace) {
			b.Locked = false
			if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
				b.Vel = rl.NewVector2(-b.Speed*0.5, -b.Speed)
			} else if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
				b.Vel = rl.NewVector2(b.Speed*0.5, -b.Speed)

			} else {
				b.Vel = rl.NewVector2(0, -b.Speed)
			}
		}
	} else {
		b.Pos.X += b.Vel.X * rl.GetFrameTime()
		b.Pos.Y += b.Vel.Y * rl.GetFrameTime()

		if b.Pos.X <= b.Radius || b.Pos.X >= screenWidth-b.Radius {
			b.Vel.X *= -1
		}
		if b.Pos.Y <= b.Radius {
			b.Vel.Y *= -1
		}

		if b.Pos.Y >= paddle.Pos.Y-b.Radius && b.Pos.X >= paddle.Pos.X && b.Pos.X <= paddle.Pos.X+paddle.Size.X {
			b.Vel.Y *= -1
			b.Vel.X = (b.Pos.X - (paddle.Pos.X + paddle.Size.X/2)) * 5
		}

		for i := range *blocks {
			if (*blocks)[i].Alive &&
				b.Pos.X >= (*blocks)[i].Pos.X && b.Pos.X <= (*blocks)[i].Pos.X+(*blocks)[i].Size.X &&
				b.Pos.Y >= (*blocks)[i].Pos.Y && b.Pos.Y <= (*blocks)[i].Pos.Y+(*blocks)[i].Size.Y {
				(*blocks)[i].Alive = false
				b.Vel.Y *= -1
				b.Speed += 20
				b.Vel = rl.Vector2Scale(rl.Vector2Normalize(b.Vel), b.Speed)
				if (len(*blocks)) == 0 {
				}
			}
		}

		if b.Pos.Y >= screenHeight {
			//b.Locked = true
			//b.Pos = rl.NewVector2(screenWidth/2, screenHeight-40)
			//b.Vel = rl.NewVector2(0, 0)

			resetGame(paddle, b, blocks)
		}
	}
}

func (b *Ball) Draw() {
	rl.DrawCircle(int32(b.Pos.X), int32(b.Pos.Y), b.Radius, rl.White)
}
