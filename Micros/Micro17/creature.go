package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Creature struct {
	Pos      rl.Vector2
	Vel      rl.Vector2
	Size     rl.Vector2
	FeetSize rl.Vector2
	Color    rl.Color
}

func (c *Creature) ApplyGravity(g rl.Vector2) {
	c.Vel = rl.Vector2Add(c.Vel, rl.Vector2Scale(g, rl.GetFrameTime()))
}

func (c *Creature) UpdateCreature() {
	c.Pos = rl.Vector2Add(c.Pos, rl.Vector2Scale(c.Vel, rl.GetFrameTime()))
}

func (c Creature) DrawCreature() {
	rl.DrawRectangle(int32(c.Pos.X), int32(c.Pos.Y), int32(c.Size.X), int32(c.Size.Y), c.Color)
}

func (c Creature) CanJump(blocks []Block) bool {
	for _, blocker := range blocks {
		if rl.CheckCollisionRecs(
			rl.NewRectangle(c.Pos.X-c.FeetSize.X, c.Pos.Y+c.Size.Y, c.FeetSize.X*2+c.Size.X, c.FeetSize.Y),
			rl.NewRectangle(blocker.Pos.X, blocker.Pos.Y, blocker.Size.X, blocker.Size.Y),
		) {
			return true
		}
	}
	return false
}

func (c *Creature) Jump(blocks []Block) {
	if c.CanJump(blocks) {
		c.Vel.Y = -350
	}
}
