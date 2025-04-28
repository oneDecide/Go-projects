package main

type DungeonGrid struct {
	cells    [][]int
	width    int32
	height   int32
	gridSize int32
}

func NewDungeonGrid(screenWidth, screenHeight, gridSize int32) *DungeonGrid {
	width := screenWidth / gridSize
	height := screenHeight / gridSize
	cells := make([][]int, width)
	for i := range cells {
		cells[i] = make([]int, height)
	}
	return &DungeonGrid{
		cells:    cells,
		width:    width,
		height:   height,
		gridSize: gridSize,
	}
}

func (g *DungeonGrid) Reset() {
	for x := range g.cells {
		for y := range g.cells[x] {
			g.cells[x][y] = 0
		}
	}
}

func (g *DungeonGrid) AddRoom(x, y, w, h int32) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			if i < g.width && j < g.height {
				g.cells[i][j] = 1
			}
		}
	}
}

func (g *DungeonGrid) AddCorridor(x1, y1, x2, y2 int32) {
	sx, ex := min(x1, x2), max(x1, x2)
	for x := sx; x <= ex; x++ {
		if x < g.width && y1 < g.height {
			g.cells[x][y1] = 2
		}
	}

	sy, ey := min(y1, y2), max(y1, y2)
	for y := sy; y <= ey; y++ {
		if x2 < g.width && y < g.height {
			g.cells[x2][y] = 2
		}
	}
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
