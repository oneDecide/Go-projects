package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Creature struct {
	Pos        rl.Vector2
	Vel        rl.Vector2
	Size       float32
	FeetSize   rl.Vector2
	Color      rl.Color
	Speed      float32
	Direction  float32
	doubleJump bool //added
	AnimationFSM
}

func (c *Creature) ApplyGravity(g rl.Vector2) {
	c.Vel = rl.Vector2Add(c.Vel, rl.Vector2Scale(g, rl.GetFrameTime()))
}

func (c *Creature) Move(x float32) {
	if x != 0 {
		c.Direction = x
	}
	c.Vel.X = x * c.Speed
}

func (c *Creature) UpdateCreature(blocks []Block) {
	c.Pos = rl.Vector2Add(c.Pos, rl.Vector2Scale(c.Vel, rl.GetFrameTime()))

	if !c.IsGrounded(blocks) || c.Vel.Y > 1 || c.Vel.Y < -1 { //if moving up or down, don't change animation
		return
	}

	if c.Vel.X == 0 {
		c.AnimationFSM.ChangeAnimationState("idle")
		return
	}

	if c.Vel.X != 0 {
		c.AnimationFSM.ChangeAnimationState("walk")
		return
	}
}

func (c *Creature) GetCanJumpRect() rl.Rectangle {
	return rl.NewRectangle(c.Pos.X-c.FeetSize.X, c.Pos.Y, c.FeetSize.X*2+c.Size, c.FeetSize.Y+c.Size)
}

func (c *Creature) GetIsGroundedRect() rl.Rectangle {
	return rl.NewRectangle(c.Pos.X, c.Pos.Y+c.Size, c.Size, c.FeetSize.Y)
}

func (c *Creature) DrawCreature() {
	// rl.DrawRectangle(int32(c.Pos.X), int32(c.Pos.Y), int32(c.Size), int32(c.Size), rl.NewColor(0, 0, 255, 64))
	// rl.DrawRectangleRec(c.GetCanJumpRect(), rl.NewColor(255, 0, 0, 64))
	// rl.DrawRectangleRec(c.GetIsGroundedRect(), rl.NewColor(0, 255, 0, 64))
	c.AnimationFSM.DrawWithFSM(c.Pos, c.Size, c.Direction)
}

func (c Creature) CanJump(blocks []Block) bool {
	for _, blocker := range blocks {
		if rl.CheckCollisionRecs(
			c.GetCanJumpRect(),
			rl.NewRectangle(blocker.Pos.X, blocker.Pos.Y, blocker.Size.X, blocker.Size.Y),
		) {
			return true
		}
	}
	return false
}

func (c Creature) IsGrounded(blocks []Block) bool {
	for _, blocker := range blocks {
		if rl.CheckCollisionRecs(
			c.GetIsGroundedRect(),
			rl.NewRectangle(blocker.Pos.X, blocker.Pos.Y, blocker.Size.X, blocker.Size.Y),
		) {
			return true
		}
	}
	return false
}

func (c *Creature) Jump(blocks []Block) {
	if c.CanJump(blocks) {
		c.doubleJump = false
		c.Vel.Y = -350
		c.CurrentAnim.Reset()
		c.ChangeAnimationState("jump")
	} else if !c.doubleJump {
		c.doubleJump = true
		c.Vel.Y = -350
		c.CurrentAnim.Reset()
		c.ChangeAnimationState("jump")
	}
}
