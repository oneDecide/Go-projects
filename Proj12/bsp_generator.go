package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BSPGenerator struct {
	Root     *BSPNode
	gridSize int32
	width    int32
	height   int32
}

func NewBSPGenerator(screenWidth, screenHeight, gridSize int32) *BSPGenerator {
	return &BSPGenerator{
		gridSize: gridSize,
		width:    screenWidth / gridSize,
		height:   screenHeight / gridSize,
	}
}

func (g *BSPGenerator) Generate() {
	g.Root = &BSPNode{
		Rect:  Rectangle{0, 0, g.width, g.height},
		Color: randomColor(),
	}
	g.splitNode(g.Root, maxDepth)
}

func (g *BSPGenerator) splitNode(node *BSPNode, depth int) {
	if depth <= 0 || (node.Rect.Width < minRoomSize*2 && node.Rect.Height < minRoomSize*2) {
		return
	}

	splitHorizontal := rand.Intn(2) == 0
	if splitHorizontal && node.Rect.Height >= minRoomSize*2 {
		splitRatio := 0.4 + rand.Float64()*0.2
		splitY := node.Rect.Y + int32(float64(node.Rect.Height)*splitRatio)

		node.Left = &BSPNode{
			Rect:  Rectangle{node.Rect.X, node.Rect.Y, node.Rect.Width, splitY - node.Rect.Y},
			Color: randomColor(),
		}
		node.Right = &BSPNode{
			Rect:  Rectangle{node.Rect.X, splitY, node.Rect.Width, node.Rect.Height - (splitY - node.Rect.Y)},
			Color: randomColor(),
		}
	} else if !splitHorizontal && node.Rect.Width >= minRoomSize*2 {
		splitRatio := 0.4 + rand.Float64()*0.2
		splitX := node.Rect.X + int32(float64(node.Rect.Width)*splitRatio)

		node.Left = &BSPNode{
			Rect:  Rectangle{node.Rect.X, node.Rect.Y, splitX - node.Rect.X, node.Rect.Height},
			Color: randomColor(),
		}
		node.Right = &BSPNode{
			Rect:  Rectangle{splitX, node.Rect.Y, node.Rect.Width - (splitX - node.Rect.X), node.Rect.Height},
			Color: randomColor(),
		}
	} else {
		return
	}

	g.splitNode(node.Left, depth-1)
	g.splitNode(node.Right, depth-1)
}

func (g *BSPGenerator) CreateRooms(grid *DungeonGrid) {
	g.createRooms(g.Root, grid)
}

func (g *BSPGenerator) createRooms(node *BSPNode, grid *DungeonGrid) {
	if node == nil {
		return
	}

	if node.Left != nil || node.Right != nil {
		g.createRooms(node.Left, grid)
		g.createRooms(node.Right, grid)
		if node.Left != nil && node.Right != nil {
			g.connectRooms(node.Left, node.Right, grid)
		}
		return
	}

	maxW := node.Rect.Width - 2
	maxH := node.Rect.Height - 2

	if maxW < minRoomSize || maxH < minRoomSize {
		return
	}

	roomW := safeRandomSize(minRoomSize, maxW, maxRoomSize)
	roomH := safeRandomSize(minRoomSize, maxH, maxRoomSize)

	maxX := node.Rect.Width - roomW
	maxY := node.Rect.Height - roomH

	if maxX <= 1 || maxY <= 1 {
		return
	}

	if (maxX-1) <= 0 || (maxY-1) <= 0 {
		return
	}

	roomX := node.Rect.X + 1 + rand.Int31n(maxX-1)
	roomY := node.Rect.Y + 1 + rand.Int31n(maxY-1)

	node.Room = Rectangle{roomX, roomY, roomW, roomH}
	grid.AddRoom(roomX, roomY, roomW, roomH)
}

func (g *BSPGenerator) connectRooms(left, right *BSPNode, grid *DungeonGrid) {
	lRoom := left.FindRoom()
	rRoom := right.FindRoom()

	if lRoom.Width == 0 || rRoom.Width == 0 {
		return
	}

	lx := lRoom.X + lRoom.Width/2
	ly := lRoom.Y + lRoom.Height/2
	rx := rRoom.X + rRoom.Width/2
	ry := rRoom.Y + rRoom.Height/2

	grid.AddCorridor(lx, ly, rx, ry)
}

func safeRandomSize(minSize, availableMax, absoluteMax int32) int32 {
	maxPossible := min(availableMax, absoluteMax)
	if maxPossible < minSize {
		return 0
	}
	validRange := maxPossible - minSize
	if validRange <= 0 {
		return minSize
	}
	return minSize + rand.Int31n(validRange)
}

func randomColor() rl.Color {
	return rl.NewColor(
		uint8(rand.Intn(200)+55),
		uint8(rand.Intn(200)+55),
		uint8(rand.Intn(200)+55),
		100,
	)
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
