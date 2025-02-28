package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var alive bool = true
	var win bool = false
	rl.InitWindow(800, 460, "2-bit Moonring")
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	playerColor := rl.SkyBlue
	playerSprite := rl.LoadTexture("textures/StickMan.png")
	enemySprite := rl.LoadTexture("textures/StickMan.png")

	var MainCharacter Character = NewCharacter(rl.NewVector2(20, 20), 0, 0, 3, 90, playerSprite, playerColor)

	enemies := make([]Character, 5)
	for i := range enemies {
		xval := rand.Intn(16)
		yval := rand.Intn(9)
		pos := rl.NewVector2(20+float32(xval)*50, 20+float32(yval)*50)
		enemies[i] = NewCharacter(pos, xval, yval, float32(rand.Intn(8)), 90, enemySprite, rl.Red)
	}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if alive && !win {
			rl.ClearBackground(rl.Gray)

			tempMove := rl.Vector2Zero()
			// Move left
			if rl.IsKeyPressed(rl.KeyA) && MainCharacter.xval > 0 {
				tempMove.X -= 1000
				MainCharacter.addBounds(-1, 0)
				fmt.Println("pos: ", MainCharacter.xval, ", ", MainCharacter.yval)
				fmt.Println("vector: ", MainCharacter.Position)
			}
			// Move right
			if rl.IsKeyPressed(rl.KeyD) && MainCharacter.xval < 15 {
				tempMove.X += 1000
				MainCharacter.addBounds(1, 0)
				fmt.Println("pos: ", MainCharacter.xval, ", ", MainCharacter.yval)
				fmt.Println("vector: ", MainCharacter.Position)
			}
			// Move up
			if rl.IsKeyPressed(rl.KeyW) && MainCharacter.yval > 0 {
				tempMove.Y -= 1000
				MainCharacter.addBounds(0, -1)
				fmt.Println("pos: ", MainCharacter.xval, ", ", MainCharacter.yval)
				fmt.Println("vector: ", MainCharacter.Position)
			}
			// Move down
			if rl.IsKeyPressed(rl.KeyS) && MainCharacter.yval < 8 {
				tempMove.Y += 1000
				MainCharacter.addBounds(0, 1)
				fmt.Println("pos: ", MainCharacter.xval, ", ", MainCharacter.yval)
				fmt.Println("vector: ", MainCharacter.Position)
			}
			MainCharacter.Move(tempMove)
			tempMove = rl.NewVector2(0, 0)
			MainCharacter.Draw()

			for i, enemy := range enemies {
				enemy.Draw()

				// Check for collision between player and enemy
				if rl.CheckCollisionRecs(
					rl.NewRectangle(MainCharacter.Position.X, MainCharacter.Position.Y, float32(MainCharacter.Sprite.Width), float32(MainCharacter.Sprite.Height)),
					rl.NewRectangle(enemy.Position.X, enemy.Position.Y, float32(enemy.Sprite.Width), float32(enemy.Sprite.Height)),
				) {
					// Compare player and enemy levels
					isPlayerStronger, enemyLevel := MainCharacter.Compare(enemy)

					if isPlayerStronger {
						// Player is stronger: defeat the enemy
						fmt.Println("Player defeated enemy!", enemyLevel)

						MainCharacter.LevelUp(enemy.level)
						enemies = append(enemies[:i], enemies[i+1:]...)
						i--
						if len(enemies) == 0 {
							win = true
						}
					} else {
						// Player is weaker: game over
						fmt.Println("Player was defeated!", enemyLevel)
						alive = false

					}
				}
			}

		}
		if win == true {
			rl.ClearBackground(rl.Green)
			MainCharacter.Position = rl.NewVector2(20, 20)
			MainCharacter.yval = 0
			MainCharacter.level = 3
			MainCharacter.SetLevel(3)
			MainCharacter.xval = 0
			rl.DrawText("Congradulations!", 0, 10, 50, rl.Gray)
			scoreText := "You Win!"
			rl.DrawText(scoreText, 0, 60, 50, rl.Gray)
			rl.DrawText("Press [ R ] to try again!", 10, 120, 40, rl.Gray)
			if rl.IsKeyDown(rl.KeyR) {
				alive = true
				win = false
				enemies = make([]Character, 5)
				for i := range enemies {
					xval := rand.Intn(16)
					yval := rand.Intn(9)
					pos := rl.NewVector2(20+float32(xval)*50, 20+float32(yval)*50)
					enemies[i] = NewCharacter(pos, xval, yval, float32(rand.Intn(8)), 90, enemySprite, rl.Red)
				}
			}
		}
		if alive == false {
			rl.ClearBackground(rl.Red)
			MainCharacter.Position = rl.NewVector2(20, 20)
			MainCharacter.yval = 0
			MainCharacter.xval = 0
			MainCharacter.level = 3
			MainCharacter.SetLevel(3)
			rl.DrawText("GAME OVER:", 0, 10, 50, rl.Gray)
			rl.DrawText("Press [ R ] to try again!", 10, 120, 40, rl.Gray)
			if rl.IsKeyDown(rl.KeyR) {
				alive = true
				enemies = make([]Character, 5)
				for i := range enemies {
					xval := rand.Intn(16)
					yval := rand.Intn(9)
					pos := rl.NewVector2(20+float32(xval)*50, 20+float32(yval)*50)
					enemies[i] = NewCharacter(pos, xval, yval, float32(rand.Intn(8)), 90, enemySprite, rl.Red)
				}
			}
		}

		rl.EndDrawing()
	}
}
