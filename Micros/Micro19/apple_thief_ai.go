package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AIState int

const (
	Seeking   = 0
	Gathering = 1
	Returning = 2
	Patrol    = 3
)

type AppleThiefAI struct {
	Creature     *Creature
	State        AIState
	SightRange   float32
	TargetPos    rl.Vector2
	ScoreZone    *ScoreZone
	worldApples  *[]*Apple
	ChangedState bool
	Timer        float32
	TickCount    int
}

func NewAppleThiefAI(creature *Creature, scoreZone *ScoreZone, worldApples *[]*Apple) *AppleThiefAI {
	return &AppleThiefAI{
		Creature:     creature,
		State:        Seeking,
		SightRange:   1000,
		ScoreZone:    scoreZone,
		worldApples:  worldApples,
		ChangedState: false,
		Timer:        0,
		TickCount:    0,
	}
}

func (ai *AppleThiefAI) SetState(newState AIState) {
	ai.ChangedState = true
	ai.State = newState
}

func (ai *AppleThiefAI) Tick() {
	if ai.ChangedState {
		ai.Timer = 0
		ai.TickCount = 0
		ai.ChangedState = false
	}

	ai.Timer += rl.GetFrameTime() 
	ai.TickCount++

	switch ai.State {
	case Seeking:
		ai.TickSeek()
	case Gathering:
		ai.TickGather()
	case Returning:
		ai.TickReturn()
	case Patrol:
		ai.TickPatrol()
	}
}

func (ai *AppleThiefAI) FindNearestApple() (*Apple, bool) {
	var nearestApple *Apple = nil
	minDist := float32(ai.SightRange)

	for _, apple := range *ai.worldApples {
		if apple.Carried {
			continue
		}
		dist := rl.Vector2Distance(ai.Creature.Pos, apple.Pos)
		if dist > ai.SightRange {
			continue
		}
		if dist < minDist {
			minDist = dist
			nearestApple = apple
		}
	}
	return nearestApple, nearestApple != nil
}

func (ai *AppleThiefAI) TickPatrol() {
	if len(ai.Creature.Apples) >= CREATURE_MAX_APPLES {
		ai.SetState(Returning)
		return
	}

	if ai.TickCount == 0 || ai.Timer >= 5.0 {
		randomOffset := rl.Vector2{
			X: float32(rl.GetRandomValue(-300, 300)),
			Y: float32(rl.GetRandomValue(-300, 300)),
		}
		ai.TargetPos = rl.Vector2Add(ai.ScoreZone.Pos, randomOffset)
		ai.Timer = 0 
	}

	if apple, found := ai.FindNearestApple(); found {
		ai.TargetPos = apple.Pos
		ai.SetState(Gathering)
		return
	}

	dist := rl.Vector2Distance(ai.Creature.Pos, ai.TargetPos)
	if dist < 10 { 
		ai.Creature.Stop()
	} else {
		ai.Creature.MoveCreatureTowards(ai.TargetPos)
	}
}

func (ai *AppleThiefAI) TickSeek() {
	if len(ai.Creature.Apples) >= CREATURE_MAX_APPLES {
		ai.SetState(Returning)
		return
	}

	if apple, found := ai.FindNearestApple(); found {
		ai.TargetPos = apple.Pos
		ai.SetState(Gathering)
	} else if len(ai.Creature.Apples) > 0 {
		ai.SetState(Returning)
	} else {
		ai.SetState(Patrol)
	}
}

func (ai *AppleThiefAI) TickGather() {
	dist := rl.Vector2Distance(ai.Creature.Pos, ai.TargetPos)

	if dist < APPLE_SIZE+CREATURE_SIZE {
		ai.Creature.GatherApples(ai.worldApples)
		ai.SetState(Seeking)
		ai.Creature.Stop()
		return
	}

	ai.Creature.MoveCreatureTowards(ai.TargetPos)
}

func (ai *AppleThiefAI) TickReturn() {
	if len(ai.Creature.Apples) == 0 {
		ai.SetState(Seeking)
		return
	}

	dist := rl.Vector2Distance(ai.Creature.Pos, ai.ScoreZone.Pos)

	if dist < SCORE_ZONE_SIZE {
		ai.Creature.DepositApple(ai.ScoreZone)
		if len(ai.Creature.Apples) == 0 {
			ai.SetState(Seeking)
		}
		ai.Creature.Stop()
		return
	}

	ai.Creature.MoveCreatureTowards(ai.ScoreZone.Pos)
}

func (ai *AppleThiefAI) TickRest() {
	if ai.TickCount == 0 {
		fmt.Println("I'm resting :3")
		ai.Creature.Stop()
	}

	if ai.Timer < 5 { //do nothing for 5 seconds
		return
	}
	ai.SetState(Seeking)
}
