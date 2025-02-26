package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Character struct {
	SpriteRenderer SpriteRenderer
	level          float32
}

func NewCharacter(pos rl.Vector2, level, size float32, sprite rl.Texture2D, color rl.Color) Character {
	sr := NewSpriteRenderer(sprite, color, pos)
	return Character{
		SpriteRenderer: sr,
		level:          level,
	}
}

func (c *Character) Move(offset rl.Vector2) {
	c.Position = rl.Vector2Add(c.Position, rl.Vector2Scale(offset, c.level*rl.GetFrameTime()))
}

func (c *Character) LevelUp(toAdd int) {
	c.level += toAdd
}
