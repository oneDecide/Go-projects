package main

import rl "github.com/gen2brain/raylib-go/raylib"

func DrawBSPPartitions(node *BSPNode) {
	if node == nil {
		return
	}

	rl.DrawRectangle(
		node.Rect.X*20,
		node.Rect.Y*20,
		node.Rect.Width*20,
		node.Rect.Height*20,
		node.Color,
	)

	DrawBSPPartitions(node.Left)
	DrawBSPPartitions(node.Right)
}

func (g *DungeonGrid) Draw() {
	for x := int32(0); x < g.width; x++ {
		for y := int32(0); y < g.height; y++ {
			switch g.cells[x][y] {
			case 1:
				rl.DrawRectangle(
					x*g.gridSize,
					y*g.gridSize,
					g.gridSize,
					g.gridSize,
					rl.Red,
				)
			case 2:
				rl.DrawRectangle(
					x*g.gridSize,
					y*g.gridSize,
					g.gridSize,
					g.gridSize,
					rl.DarkGreen,
				)
			}
		}
	}
}

func DrawDungeonGrid(gridSize, screenWidth, screenHeight int32) {
	for x := int32(0); x < screenWidth; x += gridSize {
		rl.DrawLine(x, 0, x, screenHeight, rl.Fade(rl.DarkGray, 0.2))
	}
	for y := int32(0); y < screenHeight; y += gridSize {
		rl.DrawLine(0, y, screenWidth, y, rl.Fade(rl.DarkGray, 0.2))
	}
}
