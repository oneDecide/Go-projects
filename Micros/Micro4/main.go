package main

import (
	"fmt"
	"math/rand"
)

type Creature struct {
	Name         string
	Dexterity    int
	Strength     int
	Intelligence int
}

func main() {
	var MyBasicCreature Creature = NewRPGCreature("Dave")
	fmt.Println("\nMy Basic Creature\n-----------------\nName: ", MyBasicCreature.Name)
	fmt.Println("Dex.: ", MyBasicCreature.Dexterity, "\nStr.: ", MyBasicCreature.Strength)
	fmt.Println("Int.: ", MyBasicCreature.Intelligence)
	fmt.Println()

	var MyCreature Creature = RandomRPGCreature("Arslan")
	fmt.Println("\nMy Creature\n-----------\nName: ", MyCreature.Name)
	fmt.Println("Dex.: ", MyCreature.Dexterity, "\nStr.: ", MyCreature.Strength)
	fmt.Println("Int.: ", MyCreature.Intelligence)
	fmt.Println()
}

func NewRPGCreature(named string) (newCreature Creature) {
	newCreature = Creature{Name: named, Dexterity: 1, Strength: 1, Intelligence: 1}
	return newCreature
}

func RandomRPGCreature(named string) (newCreature Creature) {
	newCreature = Creature{Name: named, Dexterity: rand.Intn(9) + 1, Strength: rand.Intn(9) + 1, Intelligence: rand.Intn(9) + 1}
	return newCreature
}
