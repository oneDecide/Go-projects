package main

import (
	"fmt"
)

func main() {
	var currentLevel int = 1
	var currentXP int = 0
	printPlayerInfo(currentLevel, currentXP)
	currentLevel, currentXP = AwardXP(currentLevel, currentXP, -126)
	printPlayerInfo(currentLevel, currentXP)
	currentLevel, currentXP = AwardXP(currentLevel, currentXP, 126)
	printPlayerInfo(currentLevel, currentXP)
	currentLevel, currentXP = AwardXP(currentLevel, currentXP, 5675)
	printPlayerInfo(currentLevel, currentXP)

}

func printPlayerInfo(currentLevel int, currentXP int) {
	fmt.Println("\nPlayer Information\n------------------")
	fmt.Println("Lvl: ", currentLevel, "\nXP: ", currentXP)
	fmt.Println()
}

func AwardXP(currentLevel int, currentXP int, awardedXP int) (newLevel int, remains int) {
	if awardedXP < 0 {
		fmt.Println("--------WARNING, Negative xp tried to be awarded, returing original values")
		return currentLevel, currentXP
	}
	var newXP int = currentXP + awardedXP
	var leftoverXP int = newXP % 100
	newLevel = currentLevel + (newXP / 100)
	return newLevel, leftoverXP
}
