package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *Game) CheckCollisions() {
	// Player-enemies
	for _, e := range g.Enemies {
		if e.IsActive && rl.CheckCollisionCircles(g.Player.Position, 20, e.Position, 15) {
			if e.EnemyType == 0 && (rl.GetTime()-float64(e.LastAttack)) > (1.0/6.0) {
				g.Player.Health -= 5
				e.LastAttack = float32(rl.GetTime())
			}
		}
	}

	// Enemy projectiles
	for i := 0; i < len(g.EnemyProjectiles); i++ {
		p := g.EnemyProjectiles[i]
		if rl.CheckCollisionCircles(g.Player.Position, 20, p.Position, 5) {
			g.Player.Health -= float32(p.Damage)
			p.Instances--
			if p.Instances <= 0 {
				g.EnemyProjectiles = append(g.EnemyProjectiles[:i], g.EnemyProjectiles[i+1:]...)
				i--
			}
		}
	}

	// Player projectiles
	for i := 0; i < len(g.PlayerProjectiles); i++ {
		p := g.PlayerProjectiles[i]
		for _, e := range g.Enemies {
			if e.IsActive && rl.CheckCollisionCircles(p.Position, 5, e.Position, 15) {
				e.TakeDamage(p.Damage)
				p.Instances--
				if p.Instances <= 0 {
					g.PlayerProjectiles = append(g.PlayerProjectiles[:i], g.PlayerProjectiles[i+1:]...)
					i--
					break
				}
			}
		}
	}

	// Diamonds
	for _, d := range g.Diamonds {
		if d.IsActive && rl.CheckCollisionCircles(g.Player.Position, 20, d.Position, 10) {
			g.Player.XP += 10
			d.IsActive = false
			g.Audio.PlayCollect()

			// Handle potential multi-level gains
			for g.Player.XP >= g.Player.NextLevelXP {
				g.LevelUpScreen = true
				g.State = "paused"
				break // Only process one level-up per frame
			}
		}
	}
}
