package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Mine struct {
	Position rl.Vector2
	radius   float32
	active   bool
}

var (
	Mines []Mine
)

func initMines() {
	Mines = make([]Mine, 0)
}

func spawnMine() {
	spawnPos := player.Position
	mine := Mine{
		Position: spawnPos,
		radius:   float32(25),
		active:   false,
	}
	Mines = append(Mines, mine)
}

func drawMine() {
	for i := range Mines {
		if !Mines[i].active {
			rl.DrawCircleV(Mines[i].Position, Mines[i].radius, rl.Black)
		} else {
			rl.DrawCircleV(Mines[i].Position, Mines[i].radius, rl.Red)
		}
	}

}

func updateMine() {
	for i := range Mines {
		if !Mines[i].active && !rl.CheckCollisionCircles(player.Position, player.BodyRadius, Mines[i].Position, Mines[i].radius) {
			Mines[i].active = true
		}
		if Mines[i].active && rl.CheckCollisionCircles(player.Position, player.BodyRadius, Mines[i].Position, Mines[i].radius) {
			GameOver()
		}
	}
}
