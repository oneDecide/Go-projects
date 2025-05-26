package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Player            *Player
	Enemies           []*Enemy
	PlayerProjectiles []*Projectile
	EnemyProjectiles  []*Projectile
	Diamonds          []*Diamond
	Audio             *AudioManager
	Camera            rl.Camera2D
	State             string
	LevelTimer        float32
	SpawnTimer        float32
	TotalGameTime     float32
	SpawnRate         float32
	LevelUpScreen     bool
	SaveLoadMessage   string
	MessageTimer      float32
	Victory           bool
}

func NewGame() *Game {
	g := &Game{
		Camera: rl.Camera2D{
			Offset:   rl.NewVector2(screenWidth/2, screenHeight/2),
			Target:   rl.NewVector2(0, 0),
			Rotation: 0,
			Zoom:     1,
		},
		State:     "menu",
		SpawnRate: 2.0,
	}

	g.Audio = NewAudioManager()
	g.ResetGame()
	return g
}

func (g *Game) ResetGame() {
	g.Player = NewPlayer()
	g.Enemies = nil
	g.PlayerProjectiles = nil
	g.EnemyProjectiles = nil
	g.Diamonds = nil
	g.SpawnTimer = 0
	g.LevelTimer = 0
	g.SpawnRate = 1.5
	g.Victory = false
	g.TotalGameTime = 0
}

func (g *Game) Update() {
	g.Audio.Update()
	frameTime := rl.GetFrameTime()

	switch g.State {
	case "menu":
		g.UpdateMenu()
	case "playing":
		g.UpdateGame(frameTime)
	case "paused":
		g.UpdatePaused()
	case "gameover":
		g.UpdateGameOver()
	case "victory":
		g.UpdateVictory()
	}

}

func (g *Game) UpdateMenu() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		g.State = "playing"
		g.Camera.Target = g.Player.Position
	}
}

func (g *Game) UpdateGame(frameTime float32) {
	// Update total game time first
	g.TotalGameTime += frameTime

	// Calculate spawn rate
	switch {
	case g.TotalGameTime < 120: // First 2 minutes
		g.SpawnRate = 1
	case g.TotalGameTime < 240: // Next 2 minutes
		g.SpawnRate = .7
	default: // Subsequent 2-minute intervals
		intervals := int((g.TotalGameTime - 240) / 120)
		g.SpawnRate = 1.0 - float32(intervals)*0.1
		if g.SpawnRate < 0.4 {
			g.SpawnRate = 0.4
		}
	}

	if g.Player.Sides >= 12 && !g.Victory {
		g.Victory = true
		g.State = "victory"
		return
	}

	if g.Player.Health <= 0 {
		g.State = "gameover"
		return
	}

	if rl.IsKeyPressed(rl.KeyEscape) {
		g.State = "paused"
	}

	if rl.IsKeyPressed(rl.KeyO) {
		g.SaveGame()
		g.SaveLoadMessage = "Game Saved!"
		g.MessageTimer = 2.0
	}
	if rl.IsKeyPressed(rl.KeyL) {
		g.LoadGame()
		g.SaveLoadMessage = "Game Loaded!"
		g.MessageTimer = 2.0
	}

	// Player movement
	var moveDir rl.Vector2
	if rl.IsKeyDown(rl.KeyW) {
		moveDir.Y -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		moveDir.Y += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		moveDir.X -= 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		moveDir.X += 1
	}
	g.Player.Move(moveDir)

	// Mouse aiming
	mouseWorldPos := rl.GetScreenToWorld2D(rl.GetMousePosition(), g.Camera)
	g.Player.Rotation = float32(math.Atan2(
		float64(mouseWorldPos.Y-g.Player.Position.Y),
		float64(mouseWorldPos.X-g.Player.Position.X),
	))

	// Shooting
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) &&
		(rl.GetTime()-float64(g.Player.LastShotTime)) > (60.0/float64(g.Player.FireRate)) {
		shootDir := rl.Vector2Normalize(rl.Vector2Subtract(mouseWorldPos, g.Player.Position))
		g.PlayerProjectiles = append(g.PlayerProjectiles, g.Player.Shoot(shootDir)...)
		g.Audio.PlayShot()
	}

	// Camera follow
	g.Camera.Target = rl.Vector2Lerp(g.Camera.Target, g.Player.Position, 0.1)

	// Update entities
	for _, p := range g.PlayerProjectiles {
		p.Update(frameTime)
	}
	for _, p := range g.EnemyProjectiles {
		p.Update(frameTime)
	}
	for _, e := range g.Enemies {
		e.Update(g.Player.Position, frameTime, g)
	}

	g.CheckCollisions()
	g.CleanupEntities()

	// Spawning
	g.SpawnTimer += frameTime
	if g.SpawnTimer >= g.SpawnRate {
		g.SpawnEnemy()
		g.SpawnTimer = 0
	}

	// Difficulty
	g.LevelTimer += frameTime
	if g.LevelTimer >= 60.0 {
		g.SpawnRate *= 0.9
		for _, e := range g.Enemies {
			e.Health += 5
			e.Speed *= 1.1
		}
		g.LevelTimer = 0
	}

	// Level up
	if g.Player.XP >= g.Player.NextLevelXP && !g.LevelUpScreen {
		g.LevelUpScreen = true
		g.State = "paused"
	}
}

func (g *Game) UpdatePaused() {
	if g.LevelUpScreen {
		if rl.IsKeyPressed(rl.KeyOne) {
			g.Player.DamageLevel++
			g.LevelUpComplete()
		}
		if rl.IsKeyPressed(rl.KeyTwo) {
			g.Player.Speed += 0.5
			g.LevelUpComplete()
		}
		if rl.IsKeyPressed(rl.KeyThree) {
			g.Player.FireRate += 30
			g.LevelUpComplete()
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.State = "playing"
		}
		if rl.IsKeyPressed(rl.KeyEscape) {
			g.State = "menu"
		}
	}
}

func (g *Game) LevelUpComplete() {
	// Increment sides and adjust XP
	g.Player.Sides++
	g.Player.NextLevelXP += 25

	// Maintain XP overflow after level up
	if g.Player.XP > g.Player.NextLevelXP {
		g.Player.XP -= g.Player.NextLevelXP - 25 // Subtract previous level's requirement
	} else {
		g.Player.XP = 0
	}

	if g.Player.Sides >= 12 {
		g.Victory = true
		g.State = "victory"
		return
	}
	// Reset level up state
	g.LevelUpScreen = false
	g.State = "playing"

	// Immediately enable new shooting angles
	g.Player.LastShotTime = 0 // Reset fire rate timer to allow instant shooting
}
func (g *Game) CleanupEntities() {
	// Cleanup enemies and spawn diamonds
	var activeEnemies []*Enemy
	for _, e := range g.Enemies {
		if e.IsActive {
			activeEnemies = append(activeEnemies, e)
		} else {
			// Spawn diamond at enemy position when dead
			g.Diamonds = append(g.Diamonds, NewDiamond(e.Position))
		}
	}
	g.Enemies = activeEnemies

	// Cleanup projectiles
	g.PlayerProjectiles = filterProjectiles(g.PlayerProjectiles)
	g.EnemyProjectiles = filterProjectiles(g.EnemyProjectiles)

	// Cleanup collected diamonds
	var activeDiamonds []*Diamond
	for _, d := range g.Diamonds {
		if d.IsActive {
			activeDiamonds = append(activeDiamonds, d)
		}
	}
	g.Diamonds = activeDiamonds
}

func filterProjectiles(projectiles []*Projectile) []*Projectile {
	var active []*Projectile
	for _, p := range projectiles {
		if !p.ShouldExpire() && p.Instances > 0 {
			active = append(active, p)
		}
	}
	return active
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Color{R: 240, G: 240, B: 240, A: 255})

	if g.State != "menu" && g.State != "gameover" {
		rl.BeginMode2D(g.Camera)
		g.DrawGrid()
		g.DrawGameWorld()
		rl.EndMode2D()
	}

	switch g.State {
	case "menu":
		g.DrawMenu()
	case "gameover":
		g.DrawGameOver()
	case "paused":
		g.DrawPauseMenu()
	case "victory":
		g.DrawVictoryScreen()
	}

	if g.State == "playing" {
		g.DrawUI()
	}

	rl.EndDrawing()
}

func (g *Game) DrawGameWorld() {
	g.Player.Draw()

	for _, p := range g.PlayerProjectiles {
		p.Draw()
	}
	for _, p := range g.EnemyProjectiles {
		p.Draw()
	}
	for _, e := range g.Enemies {
		e.Draw()
	}
	for _, d := range g.Diamonds {
		d.Draw()
	}
}

func (g *Game) DrawGrid() {
	cameraTopLeft := rl.GetScreenToWorld2D(rl.NewVector2(0, 0), g.Camera)
	cameraBottomRight := rl.GetScreenToWorld2D(rl.NewVector2(screenWidth, screenHeight), g.Camera)

	startX := float32(math.Floor(float64(cameraTopLeft.X)/gridSize)) * gridSize
	startY := float32(math.Floor(float64(cameraTopLeft.Y)/gridSize)) * gridSize
	endX := float32(math.Ceil(float64(cameraBottomRight.X)/gridSize)) * gridSize
	endY := float32(math.Ceil(float64(cameraBottomRight.Y)/gridSize)) * gridSize

	for x := startX; x <= endX; x += gridSize {
		rl.DrawLineV(rl.NewVector2(x, startY), rl.NewVector2(x, endY), rl.Color{R: 220, G: 220, B: 220, A: 255})
	}
	for y := startY; y <= endY; y += gridSize {
		rl.DrawLineV(rl.NewVector2(startX, y), rl.NewVector2(endX, y), rl.Color{R: 220, G: 220, B: 220, A: 255})
	}
}

func (g *Game) DrawMenu() {
	rl.DrawText("The Sides of Shape", screenWidth/2-150, screenHeight/2-50, 40, rl.DarkGray)
	rl.DrawText("by: KeenDLC", screenWidth/2-120, screenHeight/2-5, 18, rl.Orange)
	rl.DrawText("Press ENTER to start", screenWidth/2-130, screenHeight/2+20, 20, rl.DarkGray)
}

func (g *Game) DrawGameOver() {
	rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Fade(rl.Black, 0.7))
	rl.DrawText("GAME OVER", screenWidth/2-100, screenHeight/2-50, 40, rl.Red)
	rl.DrawText("Press ENTER to restart", screenWidth/2-120, screenHeight/2+20, 20, rl.White)
}

func (g *Game) DrawPauseMenu() {
	rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Fade(rl.Black, 0.5))
	if g.LevelUpScreen {
		rl.DrawText("LEVEL UP!", screenWidth/2-60, screenHeight/2-80, 30, rl.White)
		rl.DrawText("1. Increase Damage", screenWidth/2-80, screenHeight/2-20, 20, rl.White)
		rl.DrawText("2. Increase Speed", screenWidth/2-80, screenHeight/2+10, 20, rl.White)
		rl.DrawText("3. Increase Fire Rate", screenWidth/2-80, screenHeight/2+40, 20, rl.White)
	} else {
		rl.DrawText("PAUSED", screenWidth/2-40, screenHeight/2-20, 30, rl.White)
		rl.DrawText("Press ENTER to resume", screenWidth/2-100, screenHeight/2+20, 20, rl.White)
	}
}

func (g *Game) DrawUI() {
	// Health bar
	healthWidth := int32(200 * (g.Player.Health / 100))
	rl.DrawRectangle(10, 10, 200, 20, rl.LightGray)
	rl.DrawRectangle(10, 10, healthWidth, 20, rl.Red)

	// XP bar
	xpWidth := int32(200 * (float32(g.Player.XP) / float32(g.Player.NextLevelXP)))
	rl.DrawRectangle(10, 40, 200, 20, rl.LightGray)
	rl.DrawRectangle(10, 40, xpWidth, 20, rl.Blue)

	rl.DrawText("Press O to save, L to load", 10, 100, 20, rl.DarkGray)
	if g.MessageTimer > 0 {
		rl.DrawText(g.SaveLoadMessage, screenWidth/2-50, screenHeight-50, 20, rl.Green)
		g.MessageTimer -= rl.GetFrameTime()
	}

	// Sides count
	sidesText := fmt.Sprintf("Sides: %d", g.Player.Sides)
	rl.DrawText(sidesText, 10, 70, 20, rl.DarkGray)
}

func (g *Game) SpawnEnemy() {
	rand.Seed(time.Now().UnixNano())
	if g.SpawnRate < 0.4 {
		g.SpawnRate = 0.4
	}
	// Random angle at distance 400-600 from player
	angle := rand.Float64() * 2 * math.Pi
	distance := 400 + rand.Float64()*200
	pos := rl.NewVector2(
		g.Player.Position.X+float32(math.Cos(angle)*distance),
		g.Player.Position.Y+float32(math.Sin(angle)*distance),
	)

	// 30% chance for ranged enemy
	enemyType := 0
	if rand.Float32() < 0.3 {
		enemyType = 1
	}

	g.Enemies = append(g.Enemies, NewEnemy(pos, enemyType))
}

func (g *Game) UpdateGameOver() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		g.ResetGame()
		g.State = "playing"
	}
}

func (g *Game) Cleanup() {
	g.Audio.Cleanup()
}

func (g *Game) DrawVictoryScreen() {
	rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Fade(rl.Black, 0.7))
	rl.DrawText("VICTORY!", screenWidth/2-100, screenHeight/2-50, 50, rl.Gold)
	rl.DrawText("You've become a perfect dodecagon!", screenWidth/2-200, screenHeight/2+20, 30, rl.White)
	rl.DrawText("Press ENTER to restart", screenWidth/2-120, screenHeight/2+80, 20, rl.White)
}

func (g *Game) UpdateVictory() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		g.ResetGame()
		g.State = "playing"
	}
}
