package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Name      string
	X         float32
	Y         float32
	Inventory []*Item
	Texture   rl.Texture2D
	Color     rl.Color
}

func (c *Creature) DrawCreature() {
	DrawTextureEz(c.Texture, rl.NewVector2(c.X, c.Y), 0, 10, c.Color)
}

func (c *Creature) AddItem(item Item) {
	c.Inventory = append(c.Inventory, &item)
}

func NewCreature(newName string, newX, newY float32, sprite rl.Texture2D, newColor rl.Color) Creature {
	c := Creature{Name: newName, X: newX, Y: newY, Texture: sprite, Color: newColor}
	c.Inventory = make([]*Item, 0)
	return c
}

func DrawTextureEz(texture rl.Texture2D, pos rl.Vector2, angle float32, scale float32, color rl.Color) {
	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	destRect := rl.NewRectangle(pos.X, pos.Y, float32(texture.Width)*scale, float32(texture.Height)*scale)

	origin := rl.Vector2Scale(rl.NewVector2(float32(texture.Width)/2, float32(texture.Height)/2), scale)

	rl.DrawTexturePro(texture, sourceRect,
		destRect,
		origin, angle, color)
}

func (c *Creature) RandomizeColor() {
	c.Color = rl.NewColor(
		uint8(rand.IntN(255)),
		uint8(rand.IntN(255)),
		uint8(rand.IntN(255)),
		255,
	)
}

func (c *Creature) Save() error {
	data, err := json.MarshalIndent(c, "", "  ")
	fmt.Println("SAVING...")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return os.WriteFile(c.Name+".json", data, 0644)
}

func (c *Creature) Load() error {
	data, err := os.ReadFile(c.Name + ".json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, c)
}
