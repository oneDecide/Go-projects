package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1280
	screenHeight = 720
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
