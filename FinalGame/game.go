package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameState int

const (
	StateMainMenu GameState = iota
	StatePlaying
	StateUpgradeMenu
	StateQuitMenu
	StateGameOver
)

type Game struct {
	state       GameState
	player      *Player
	enemies     []Enemy
	diamonds    []*Diamond
	projectiles []*Projectile
	audio       *AudioManager
	camera      *Camera2DWrapper
	spawnTimer  float32
	levelTimer  float32
}

func NewGame() *Game {
	return &Game{
		state:  StateMainMenu,
		player: NewPlayer(),
		audio:  NewAudioManager(),
		camera: NewCamera2DWrapper(),
	}
}

func (g *Game) Close() {
	g.audio.Close()
}

func (g *Game) Update() {
	switch g.state {
	case StateMainMenu:
		UpdateMainMenu(g)

	case StatePlaying:
		if g.player.Health <= 0 {
			g.state = StateGameOver
			return
		}
		if rl.IsKeyPressed(rl.KeyEscape) {
			g.state = StateQuitMenu
			return
		}

		dt := rl.GetFrameTime()
		g.spawnTimer += dt
		g.levelTimer += dt
		if g.spawnTimer > EnemySpawnRate {
			g.spawnEnemies()
			g.spawnTimer = 0
		}
		if g.levelTimer > EnemyLevelUpRate {
			g.levelUpEnemies()
			g.levelTimer = 0
		}

		g.player.Update()
		g.camera.Update(g.player.Pos)
		g.updateProjectiles()
		g.updateEnemies()
		g.updateDiamonds()
		g.audio.Update()

		if g.player.XP >= g.player.XPToNext {
			g.state = StateUpgradeMenu
		}

	case StateUpgradeMenu:
		if rl.IsKeyPressed(rl.KeyOne) || rl.IsKeyPressed(rl.KeyTwo) || rl.IsKeyPressed(rl.KeyThree) {
			key := rl.GetKeyPressed()
			g.player.ApplyUpgrade(key)
			g.player.LevelUpDone()
			g.state = StatePlaying
		}

	case StateQuitMenu:
		UpdateQuitMenu(g)

	case StateGameOver:
		if rl.IsKeyPressed(rl.KeyEnter) {
			gameRef = NewGame()
		}
		if rl.IsKeyPressed(rl.KeyQ) {
			rl.CloseWindow()
		}
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.LightGray)
	rl.BeginMode2D(g.camera.Camera)

	DrawGrid(g.camera)

	switch g.state {
	case StateMainMenu:
		DrawMainMenu()

	case StatePlaying:
		g.player.Draw()
		for _, e := range g.enemies {
			e.Draw()
		}
		for _, d := range g.diamonds {
			d.Draw()
		}
		for _, p := range g.projectiles {
			p.Draw()
		}

	case StateUpgradeMenu:
		DrawUpgradeMenu()

	case StateQuitMenu:
		DrawQuitMenu()

	case StateGameOver:
		rl.DrawText("GAME OVER", ScreenWidth/2-100, ScreenHeight/2-40, 30, rl.Red)
		rl.DrawText("Press ENTER to Restart or Q to Quit", ScreenWidth/2-220, ScreenHeight/2, 20, rl.Black)
	}

	rl.EndMode2D()
	rl.EndDrawing()
}

func (g *Game) spawnEnemies() {
	g.enemies = append(g.enemies, NewCircleEnemy(g.player), NewArrowEnemy(g.player))
}
func (g *Game) levelUpEnemies() {
	for _, e := range g.enemies {
		e.LevelUp()
	}
}

func (g *Game) updateProjectiles() {
	var live []*Projectile
	for _, p := range g.projectiles {
		if !p.Update() {
			continue
		}
		hit := false
		if p.Owner == OwnerPlayer {
			for i, e := range g.enemies {
				if p.Collides(e) {
					dead := e.Hit(p.DamageInstance())
					if dead {
						g.diamonds = append(g.diamonds, NewDiamond(e.GetPosition()))
						g.player.XP += 10
						g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
					}
					if p.ConsumeInstance() {
						hit = true
						break
					}
				}
			}
		} else if p.CollidesPlayer(g.player) {
			g.player.Health -= p.Damage
			hit = true
		}
		if !hit {
			live = append(live, p)
		}
	}
	g.projectiles = live
}

func (g *Game) updateEnemies() {
	for _, e := range g.enemies {
		e.Update()
		if ce, ok := e.(*CircleEnemy); ok && ce.CollidesPlayer(g.player) {
			g.player.Health -= int(MeleeDPS * rl.GetFrameTime())
		}
	}
}

func (g *Game) updateDiamonds() {
	var live []*Diamond
	for _, d := range g.diamonds {
		if d.CollidesPlayer(g.player) {
			continue
		}
		live = append(live, d)
	}
	g.diamonds = live
}
