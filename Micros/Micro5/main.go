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

	var rosters []Creature = CreateRPGRoster(100)
	fmt.Print("\nTesting: ", rosters[0].Name)
}

func NewRPGCreature(named string) (newCreature Creature) {
	newCreature = Creature{Name: named, Dexterity: 1, Strength: 1, Intelligence: 1}
	return newCreature
}

func RandomRPGCreature(named string) (newCreature Creature) {
	newCreature = Creature{Name: named, Dexterity: rand.Intn(9) + 1, Strength: rand.Intn(9) + 1, Intelligence: rand.Intn(9) + 1}
	return newCreature
}

func CreateRPGRoster(size int) []Creature {
	var critters []Creature = make([]Creature, 0, size)
	var strCount int = 0
	var dexCount int = 0
	var intCount int = 0
	for i := 0; i < size; i++ {
		critters = append(critters, RandomRPGCreature("James"))
		strCount += critters[i].Strength
		dexCount += critters[i].Dexterity
		intCount += critters[i].Intelligence
	}
	fmt.Print("CREATED ROSTER\n---------------\n")
	fmt.Println("\nTotal Strength: ", strCount)
	fmt.Println("Total Dexterity: ", dexCount)
	fmt.Println("Total Intelligence: ", intCount)
	fmt.Println()
	return critters
}
