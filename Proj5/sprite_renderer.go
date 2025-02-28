package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// contains all the info we need for drawing a picture
type SpriteRenderer struct {
	Sprite   rl.Texture2D
	Color    rl.Color
	Position rl.Vector2
	Angle    float32
	Scale    float32
	Level    float32
}

func NewSpriteRenderer(newSprite rl.Texture2D, newColor rl.Color, newPosition rl.Vector2, level float32) SpriteRenderer {
	sr := SpriteRenderer{
		Sprite:   newSprite,
		Color:    newColor,
		Position: newPosition,
		Angle:    0,
		Scale:    2,
		Level:    level,
	}
	return sr
}
func (sr *SpriteRenderer) SetLevel(newLevel float32) {
	sr.Level = newLevel
}

// our ez draw method, using data from a sprite renderer this time
func (sr SpriteRenderer) Draw() {
	sourceRect := rl.NewRectangle(0, 0, float32(sr.Sprite.Width), float32(sr.Sprite.Height))
	destRect := rl.NewRectangle(sr.Position.X, sr.Position.Y, float32(sr.Sprite.Width)*sr.Scale, float32(sr.Sprite.Height)*sr.Scale)
	origin := rl.Vector2Scale(rl.NewVector2(float32(sr.Sprite.Width)/2, float32(sr.Sprite.Height)/2), sr.Scale)
	rl.DrawTexturePro(sr.Sprite, sourceRect,
		destRect,
		origin, sr.Angle, sr.Color)

	levelText := fmt.Sprintf("  %.0f", sr.Level)                                     // Format level to one decimal place
	textX := int32(sr.Position.X) - int32(float32(sr.Sprite.Width)*sr.Scale/2)       // Bottom-left X
	textY := int32(sr.Position.Y) + int32(float32(sr.Sprite.Height)*sr.Scale/2) - 20 // Bottom-left Y
	rl.DrawText(levelText, textX, textY, 20, rl.White)
}
