package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SCORE_ZONE_SIZE = 100
)

type ScoreZone struct {
	Pos    rl.Vector2
	Points int
}

func (sz ScoreZone) DrawScoreZone() {
	rl.DrawCircle(int32(sz.Pos.X), int32(sz.Pos.Y), SCORE_ZONE_SIZE, rl.NewColor(33, 33, 33, 255))
}

func (sz *ScoreZone) EarnPoint(Apple *Apple) {
	Apple.Pos = rl.NewVector2(1000000, 1000000)
	sz.Points += 1
}
