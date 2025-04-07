package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CREATURE_SPEED      = 150
	CREATURE_SIZE       = 64
	CREATURE_MAX_APPLES = 5
)

type Creature struct {
	Pos       rl.Vector2
	Color     rl.Color
	Apples    []*Apple
	Direction float32
	AnimationFSM
}

func (c *Creature) DrawCreature() {
	//rl.DrawCircle(int32(c.Pos.X), int32(c.Pos.Y), CREATURE_SIZE, c.Color)
	for i, a := range c.Apples {
		a.Pos = rl.NewVector2(c.Pos.X, c.Pos.Y-APPLE_SIZE-float32(APPLE_SIZE*2)*float32(i+1))
	}
	c.AnimationFSM.CurrentAnim.UpdateTime()
	c.AnimationFSM.DrawWithFSM(rl.Vector2Subtract(c.Pos, rl.NewVector2(CREATURE_SIZE/2, CREATURE_SIZE/2)), CREATURE_SIZE, c.Direction)
}

func NewCreature(Color rl.Color, animationFSM AnimationFSM) Creature {
	return Creature{
		Pos:          rl.NewVector2(float32(rl.GetScreenWidth())/2, float32(rl.GetScreenHeight())/2),
		Color:        Color,
		Apples:       make([]*Apple, 0),
		Direction:    1,
		AnimationFSM: animationFSM,
	}
}

func (c *Creature) MoveCreatureTowards(target rl.Vector2) {
	direction := rl.Vector2Subtract(target, c.Pos)
	direction = rl.Vector2Normalize(direction)
	c.MoveCreature(direction)
}

func (c *Creature) Stop() {
	c.AnimationFSM.ChangeAnimationState("idle")
}

func (c *Creature) MoveCreature(vec rl.Vector2) {
	vec = rl.Vector2Normalize(vec)
	c.Pos = rl.Vector2Add(c.Pos, rl.Vector2Scale(vec, rl.GetFrameTime()*CREATURE_SPEED))
	if vec.X != 0 || vec.Y != 0 {
		c.AnimationFSM.ChangeAnimationState("walk")
	} else {
		c.AnimationFSM.ChangeAnimationState("idle")
	}
	if vec.X > 0 {
		c.Direction = 1
	}
	if vec.X < 0 {
		c.Direction = -1
	}
}

func (c *Creature) Pickup(Apple *Apple) {
	if Apple.Carried {
		return
	}
	c.Apples = append(c.Apples, Apple)
	Apple.Carried = true
}
func (c *Creature) GatherApples(worldApples *[]*Apple) {
	for _, a := range *worldApples {
		if len(c.Apples) >= CREATURE_MAX_APPLES {
			return
		}
		if rl.Vector2Distance(c.Pos, a.Pos) < APPLE_SIZE+CREATURE_SIZE {
			c.Pickup(a)
		}
	}
}

func (c *Creature) DepositApple(scoreZone *ScoreZone) {
	if len(c.Apples) < 1 {
		return
	}
	if rl.Vector2Distance(c.Pos, scoreZone.Pos) < SCORE_ZONE_SIZE {
		scoreZone.EarnPoint(c.Apples[len(c.Apples)-1])
		c.Apples = c.Apples[:len(c.Apples)-1]
	}
}

func (c *Creature) Teleport(Pos rl.Vector2) {
	c.Pos = Pos
}
