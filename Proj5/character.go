package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Character struct {
	SpriteRenderer
	level float32
	xval  int
	yval  int
}

func NewCharacter(pos rl.Vector2, xval int, yval int, level float32, size float32, sprite rl.Texture2D, color rl.Color) Character {
	sr := NewSpriteRenderer(sprite, color, pos, level)
	return Character{
		SpriteRenderer: sr,
		level:          level,
		xval:           xval,
		yval:           yval,
	}
}

func (c *Character) Move(offset rl.Vector2) {
	c.Position = rl.Vector2Add(c.Position, rl.Vector2Scale(offset, 0.0501))
}

func (c *Character) LevelUp(toAdd float32) {
	c.level += toAdd
	c.SetLevel(c.level)
}

func (c *Character) Compare(enemy Character) (bool, int) {
	playerLevel := int(c.level)
	enemyLevel := int(enemy.level)
	fmt.Println("Player level: ", playerLevel, "  Enemy Level: ", enemyLevel)
	return playerLevel >= enemyLevel, enemyLevel
}

func (c *Character) addBounds(xval int, yval int) {
	c.xval += xval
	c.yval += yval
}
