package main

import (
	"fmt"
	"math/rand/v2"
)

type Potion struct {
	Effect     string
	Power      int
	Uses       int
	Color      string
	SideEffect string
}

func (p *Potion) Generate(seed1, seed2 uint64) {
	combinedSeed := seed1 ^ seed2
	r := rand.New(rand.NewPCG(combinedSeed, 0))

	effects := []string{
		"Invisibility", "Healing", "Strength Boost", "Agility Enhancement",
		"Night Vision", "Fire Resistance", "Water Breathing", "Levitation",
		"Mind Reading", "Luck", "Time Slow", "Telekinesis",
		"Chameleon Skin", "Regeneration", "Elemental Control", "Animal Speech",
	}
	p.Effect = effects[r.IntN(len(effects))]

	p.Power = r.IntN(100) + 1

	p.Uses = r.IntN(10) + 1

	colors := []string{
		"Red", "Blue", "Green", "Purple", "Gold", "Silver", "Black", "Rainbow",
		"Glowing", "Clear", "Swirling", "Metallic", "Iridescent",
	}
	p.Color = colors[r.IntN(len(colors))]

	sideEffects := []string{
		"None", "Temporary Blindness", "Hiccups", "Color Change",
		"Sleepiness", "Increased Appetite", "Random Teleportation",
		"Truth Compulsion", "Musical Speech", "Memory Loss",
		"Spontaneous Dancing", "Elemental Aura", "Time Perception Shift",
	}
	p.SideEffect = sideEffects[r.IntN(len(sideEffects))]
}

func main() {
	var potion Potion
	potion.Generate(rand.Uint64(), rand.Uint64())

	fmt.Println("Generated Potion:")
	fmt.Printf("Effect: %s\n", potion.Effect)
	fmt.Printf("Power: %d\n", potion.Power)
	fmt.Printf("Uses: %d\n", potion.Uses)
	fmt.Printf("Color: %s\n", potion.Color)
	fmt.Printf("Side Effect: %s\n", potion.SideEffect)
}
