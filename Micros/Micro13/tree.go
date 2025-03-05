package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Tree struct {
	Pos rl.Vector2
}

func NewTree(newPos rl.Vector2) Tree {
	newTree := Tree{Pos: newPos}
	return newTree
}

func (t Tree) DrawTree() {
	rl.DrawCircle(int32(t.Pos.X), int32(t.Pos.Y), 30, rl.Lime)
}
