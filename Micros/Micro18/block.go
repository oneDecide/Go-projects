package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Block struct {
	Pos   rl.Vector2
	Size  rl.Vector2
	Color rl.Color
}

func (bl Block) DrawBlock() {
	rl.DrawRectangle(int32(bl.Pos.X), int32(bl.Pos.Y), int32(bl.Size.X), int32(bl.Size.Y), bl.Color)
}

// func CheckPlatformCollision(creature *Creature, block Block) {
// 	if rl.CheckCollisionRecs( //Raylib let's us quickly check overlap with the rectangle class.
// 		rl.NewRectangle(creature.Pos.X, creature.Pos.Y, creature.Size, creature.Size),
// 		rl.NewRectangle(block.Pos.X, block.Pos.Y, block.Size.X, block.Size.Y),
// 	) {
// 		if creature.Pos.Y+creature.Size > block.Pos.Y && creature.Vel.Y > 0 { //now check which side to stop the velocity
// 			creature.Pos.Y = block.Pos.Y - creature.Size //move box in case of overlap
// 			creature.Vel.Y = 0                           //stop the box from moving further
// 		} else if creature.Pos.Y < block.Pos.Y+block.Size.Y && creature.Vel.Y < 0 {
// 			creature.Pos.Y = block.Pos.Y + block.Size.Y
// 			creature.Vel.Y = 0
// 		} else if creature.Pos.X+creature.Size > block.Pos.X && creature.Vel.X > 0 {
// 			creature.Pos.X = block.Pos.X - creature.Size
// 			creature.Vel.X = 0
// 		} else if creature.Pos.X < block.Pos.X+block.Size.X && creature.Vel.X < 0 {
// 			creature.Pos.X = block.Pos.X + block.Size.X
// 			creature.Vel.X = 0
// 		}
// 	}
// }

func CheckPlatformCollision(creature *Creature, block Block) {
	if rl.CheckCollisionRecs(
		rl.NewRectangle(creature.Pos.X, creature.Pos.Y, creature.Size, creature.Size),
		rl.NewRectangle(block.Pos.X, block.Pos.Y, block.Size.X, block.Size.Y),
	) {

		//resolve min overlap
		creatureRight := creature.Pos.X + creature.Size
		blockRight := block.Pos.X + block.Size.X
		overlapX := min(creatureRight, blockRight) - max(creature.Pos.X, block.Pos.X)

		creatureBottom := creature.Pos.Y + creature.Size
		blockBottom := block.Pos.Y + block.Size.Y
		overlapY := min(creatureBottom, blockBottom) - max(creature.Pos.Y, block.Pos.Y)

		if overlapX < overlapY {

			if creature.Pos.X < block.Pos.X { //creature to right of block
				creature.Pos.X = block.Pos.X - creature.Size
			} else { //creature to left of block
				creature.Pos.X = block.Pos.X + block.Size.X
			}
			creature.Vel.X = 0
		} else {

			if creature.Pos.Y < block.Pos.Y { //creature on top of block
				creature.Pos.Y = block.Pos.Y - creature.Size
				if creature.Vel.Y > 0 {
					creature.Vel.Y = 0
				}
			} else { //creature on bottom of block
				creature.Pos.Y = block.Pos.Y + block.Size.Y
				creature.Vel.Y = 0
			}
		}
	}
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
