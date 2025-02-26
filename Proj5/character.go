package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Character struct {
	SpriteRenderer
	level float32
	xval  int
	yval  int
}

func NewCharacter(pos rl.Vector2, xval int, yval int, level, size float32, sprite rl.Texture2D, color rl.Color) Character {
	sr := NewSpriteRenderer(sprite, color, pos)
	return Character{
		SpriteRenderer: sr,
		level:          level,
		xval:           xval,
		yval:           yval,
	}
}

func (c *Character) Move(offset rl.Vector2) {
	c.Position = rl.Vector2Add(c.Position, rl.Vector2Scale(offset, c.level*rl.GetFrameTime()))
}

func (c *Character) LevelUp(toAdd float32) {
	c.level += toAdd
}

func (c *Character) Compare(enemy Character) {

}

func (c *Character) addBounds(xval int, yval int) {
	c.xval += xval
	c.yval += yval
}
