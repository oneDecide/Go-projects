package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Color rl.Color
	Alive bool
}

func NewBlock(x, y float32, color rl.Color) Block {
	return Block{
		Pos:   rl.NewVector2(x, y),
		Size:  rl.NewVector2(90, 40),
		Color: color,
		Alive: true,
	}
}

func (b *Block) Draw() {
	if b.Alive {
		rl.DrawRectangleV(b.Pos, b.Size, b.Color)
	}
}
