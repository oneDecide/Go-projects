package main

import (
	"encoding/json"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	saveFile     = "breakout_save.json"
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Breakout")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	paddle := NewPaddle()
	ball := NewBall()

	blocks := make([]Block, 0)
	for i := 0; i < 3; i++ {
		for u := 0; u < 10; u++ {
			blocks = append(blocks, NewBlock(float32(100*u+125), float32(60*i+40), rl.Red))
		}
	}

	for !rl.WindowShouldClose() {
		paddle.Update()
		ball.Update(&paddle, &blocks)

		//Save game state when 'S' is pressed
		if rl.IsKeyPressed(rl.KeyS) {
			err := saveGame(paddle, ball, blocks)
			if err != nil {
				fmt.Println("Failed to save game:", err)
			} else {
				fmt.Println("Game saved!")
			}
		}

		//Load game state when 'P' is pressed
		if rl.IsKeyPressed(rl.KeyP) {
			err := loadGame(&paddle, &ball, &blocks)
			if err != nil {
				fmt.Println("Failed to load game:", err)
			} else {
				fmt.Println("Game loaded!")
			}
		}

		allBlocksDestroyed := true
		for _, block := range blocks {
			if block.Alive {
				allBlocksDestroyed = false
				break
			}
		}

		if allBlocksDestroyed {
			resetGame(&paddle, &ball, &blocks)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		paddle.Draw()
		ball.Draw()
		for _, block := range blocks {
			block.Draw()
		}

		rl.EndDrawing()
	}
}

func resetGame(paddle *Paddle, ball *Ball, blocks *[]Block) {
	*paddle = NewPaddle()
	*ball = NewBall()
	*blocks = make([]Block, 0)
	for i := 0; i < 3; i++ {
		for u := 0; u < 10; u++ {
			*blocks = append(*blocks, NewBlock(float32(100*u+125), float32(60*i+40), rl.Red))
		}
	}
}

func saveGame(paddle Paddle, ball Ball, blocks []Block) error {
	fmt.Println("SAVING...")
	data, err := json.MarshalIndent(struct {
		Paddle Paddle
		Ball   Ball
		Blocks []Block
	}{
		Paddle: paddle,
		Ball:   ball,
		Blocks: blocks,
	}, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(saveFile, data, 0644)
}

func loadGame(paddle *Paddle, ball *Ball, blocks *[]Block) error {
	data, err := os.ReadFile(saveFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &struct {
		Paddle *Paddle
		Ball   *Ball
		Blocks *[]Block
	}{
		Paddle: paddle,
		Ball:   ball,
		Blocks: blocks,
	})
}
