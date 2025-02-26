package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var alive bool = true

	rl.InitWindow(800, 460, "2-bit Moonring")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	playerColor := rl.SkyBlue
	playerSprite := rl.LoadTexture("textures/StickMan.png")

	//EnemyCharacter := rl.LoadTexture("textures/StickMan.png")
	var MainCharacter Character = NewCharacter(rl.NewVector2(20, 20), 0, 0, 3, 90, playerSprite, playerColor)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if alive {
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

		}
		if alive == false {
			rl.ClearBackground(rl.Red)
			rl.DrawText("GAME OVER:", 0, 10, 50, rl.Gray)
			scoreText := "Score: "
			rl.DrawText(scoreText, 0, 60, 50, rl.Gray)
			rl.DrawText("Press [ R ] to try again!", 10, 120, 40, rl.Gray)
			if rl.IsKeyDown(rl.KeyR) {
				alive = true
			}
		}

		rl.EndDrawing()
	}
}
