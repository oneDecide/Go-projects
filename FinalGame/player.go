package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "math"
)

const (
    OwnerPlayer = 0
    OwnerEnemy  = 1
)

type Player struct {
    Pos         rl.Vector2
    Speed       float32
    Angle       float32
    Health      int
    XP          int
    XPToNext    int
    Level       int
    Sides       int
    FireRate    float32
    DamageLevel int
    lastShot    float64      // now matches rl.GetTime() (float64)
}

func NewPlayer() *Player {
    return &Player{
        Pos:         rl.Vector2{X: 0, Y: 0},
        Speed:       200,
        Health:      100,
        XPToNext:    100,
        Level:       1,
        Sides:       1,
        FireRate:    2.0,
        DamageLevel: 1,
        lastShot:    0,
    }
}

func (p *Player) Update() {
    dt := rl.GetFrameTime()
    dir := rl.Vector2{X: 0, Y: 0}
    if rl.IsKeyDown(rl.KeyW) { dir.Y-- }
    if rl.IsKeyDown(rl.KeyS) { dir.Y++ }
    if rl.IsKeyDown(rl.KeyA) { dir.X-- }
    if rl.IsKeyDown(rl.KeyD) { dir.X++ }
    if rl.Vector2Length(dir) > 0 {
        dir = rl.Vector2Normalize(dir)
        p.Pos = rl.Vector2Add(p.Pos, rl.Vector2Scale(dir, p.Speed*dt))
    }

    // Convert mouse to world space for correct aiming
    worldMouse := rl.GetScreenToWorld2D(rl.GetMousePosition(), gameRef.camera.Camera)
    delta := rl.Vector2Subtract(worldMouse, p.Pos)
    p.Angle = float32(math.Atan2(float64(delta.Y), float64(delta.X)))

    now := rl.GetTime()
    if rl.IsMouseButtonDown(rl.MouseLeftButton) && now-p.lastShot > 1.0/float64(p.FireRate) {
        p.lastShot = now
        for _, pt := range p.getVertices() {
            dirN := rl.Vector2Normalize(rl.Vector2Subtract(pt, p.Pos))
            proj := NewProjectile(pt, dirN, 5, OwnerPlayer)
            proj.Instances = p.DamageLevel
            gameRef.projectiles = append(gameRef.projectiles, proj)
        }
    }
}

func (p *Player) getVertices() []rl.Vector2 {
	var pts []rl.Vector2
	r := float32(20)
	switch p.Sides {
	case 1:
		front := p.Angle
		back := p.Angle + math.Pi
		pts = []rl.Vector2{
			{X: p.Pos.X + r*float32(math.Cos(float64(front))), Y: p.Pos.Y + r*float32(math.Sin(float64(front)))},
			{X: p.Pos.X + r*float32(math.Cos(float64(back))), Y: p.Pos.Y + r*float32(math.Sin(float64(back)))},
		}
	case 2:
		front := p.Angle
		back := p.Angle + math.Pi
		pts = []rl.Vector2{
			{X: p.Pos.X + r*float32(math.Cos(float64(front))), Y: p.Pos.Y + r*float32(math.Sin(float64(front)))},
			{X: p.Pos.X + r*float32(math.Cos(float64(back))), Y: p.Pos.Y + r*float32(math.Sin(float64(back)))},
		}
	default:
		for i := 0; i < p.Sides; i++ {
			ang := float64(p.Angle) + 2*math.Pi*float64(i)/float64(p.Sides)
			pts = append(pts, rl.Vector2{
				X: p.Pos.X + r*float32(math.Cos(ang)),
				Y: p.Pos.Y + r*float32(math.Sin(ang)),
			})
		}
	}
	return pts
}

func (p *Player) ApplyUpgrade(key int32) {
	switch key {
	case rl.KeyOne:
		p.DamageLevel++
	case rl.KeyTwo:
		p.Speed += 50
	case rl.KeyThree:
		p.FireRate += 0.5
	}
}

func (p *Player) LevelUpDone() {
	p.Level++
	p.XP -= p.XPToNext
	p.XPToNext += 25
	p.Sides++
}

func (p *Player) Draw() {
	verts := p.getVertices()
	for i := range verts {
		next := verts[(i+1)%len(verts)]
		rl.DrawLineV(verts[i], next, rl.Blue)
	}
	// Health bar
	rl.DrawRectangle(int32(p.Pos.X-20), int32(p.Pos.Y-30), 40, 5, rl.Red)
	rl.DrawRectangle(int32(p.Pos.X-20), int32(p.Pos.Y-30), int32(40*float32(p.Health)/100), 5, rl.Green)
	// XP bar
	rl.DrawRectangle(int32(p.Pos.X-20), int32(p.Pos.Y-24), 40, 5, rl.DarkGray)
	rl.DrawRectangle(int32(p.Pos.X-20), int32(p.Pos.Y-24), int32(40*float32(p.XP)/float32(p.XPToNext)), 5, rl.Yellow)
}
