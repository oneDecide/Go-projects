package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Creature struct {
	SpriteRenderer
	speed float32
}

func NewCreature(pos rl.Vector2, speed, size float32, sprite rl.Texture2D, color rl.Color) Creature {
	sr := NewSpriteRenderer(sprite, color, pos)
	return Creature{
		SpriteRenderer: sr,
		speed:          speed,
	}
}

func (c *Creature) Move(offset rl.Vector2) {
	c.Position = rl.Vector2Add(c.Position, rl.Vector2Scale(offset, c.speed*rl.GetFrameTime()))
}
