package main

//import rl "github.com/gen2brain/raylib-go/raylib"

type Health struct {
	maxHealth     float32
	currentHealth float32
}

func NewHealth(maxHealthNew float32) Health {
	return Health{
		maxHealth:     maxHealthNew,
		currentHealth: maxHealthNew,
	}
}

func (h *Health) Damage(takeDamage float32) {
	h.currentHealth -= takeDamage
}

func (h *Health) Heal(toHeal float32) {
	h.currentHealth += toHeal
	if h.currentHealth > h.maxHealth {
		h.currentHealth = h.maxHealth
	}
}
