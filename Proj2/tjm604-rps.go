package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("\n\nTHE GAME IS STARTING\n")
	fmt.Print("This is a game of Rock Paper Scissors with you and your AI enemy having 10 lives.\n")
	fmt.Print("The Damage number increases by 1 right before every match and resets to 0 when damage is dealt\n")
	fmt.Print("If there is a draw, damage increases and is brought over into the next round.\n")
	fmt.Print("------------------------------------------------------------------------------------------------\n\n\n")

	var damage int = 0
	var rounds int = 0
	var PlayerHealth int = 10
	var AIHealth int = 10
	var playerChoice int = -1
	for {
		rounds++
		damage++
		fmt.Println("NEW GAME START\n----------------------")
		fmt.Print("Your Health: ", PlayerHealth, " /// AI Health: ", AIHealth, "\nDamage this round: ", damage, "\n\n")
		fmt.Print("Round - ", rounds, "\n\n")
		fmt.Print("Inputs: (1) = Rock, (2) = Paper, (3) = Scissors (4) Random *2X DAMAGE*\nYour Input - ")
		fmt.Scan(&playerChoice)
		fmt.Println()
		playerChoice, damage = playerChose(playerChoice, damage)
		fmt.Println()

		damage, PlayerHealth, AIHealth = RPSGAME(damage, PlayerHealth, AIHealth, playerChoice)
		time.Sleep(time.Second * 2)
		if PlayerHealth < 1 {
			fmt.Println("FINISH! - YOU LOSE")
			break
		}
		if AIHealth < 1 {
			fmt.Println("FINISH! - YOU WIN!")
			break
		}
	}
}

func playerChose(playerChoice int, damage int) (choiceRet int, damageRet int) {
	damageRet = damage
	choiceRet = playerChoice
	switch playerChoice {
	case 1:
		fmt.Print("You Chose Rock")
	case 2:
		fmt.Print("You Chose Paper")
	case 3:
		fmt.Print("You Chose Scissors")
	case 4:
		fmt.Print("You Chose Random, The Damage has doubled\n")
		damageRet *= 2
		fmt.Print("New damage: ", damageRet, "\n\n")
		choiceRet = rand.Intn(3) + 1
		choiceRet, damageRet = playerChose(choiceRet, damageRet)
	default:
		fmt.Print("\nIncorrect Input, try again: ")
		fmt.Scan(&choiceRet)
		choiceRet, damageRet = playerChose(choiceRet, damageRet)
	}
	return choiceRet, damageRet
}

func RPSGAME(damage int, playerH int, AIH int, playerChoice int) (damRet int, NewPHealth int, NewAIHealth int) {
	var AIchoice int = rand.Intn(3) + 1
	NewPHealth = playerH
	NewAIHealth = AIH
	damRet = damage
	fmt.Print("THE AI CHOSE - ")
	switch AIchoice {
	case 1:
		fmt.Println("ROCK")
	case 2:
		fmt.Println("PAPER")
	case 3:
		fmt.Println("SCISSORS")
	}
	if AIchoice == playerChoice {
		fmt.Print("\nTHERE IS A DRAW!!!!!\n\n")
		return damRet, NewPHealth, NewAIHealth
	}
	if (playerChoice == 1 && AIchoice == 2) || (playerChoice == 2 && AIchoice == 3) || (playerChoice == 3 && AIchoice == 1) {
		fmt.Println("\nTHE AI WON THIS ROUND - DAMAGE: ", damRet)
		NewPHealth -= damRet
		damRet = 0
		return damRet, NewPHealth, NewAIHealth
	}
	if (playerChoice == 1 && AIchoice == 3) || (playerChoice == 2 && AIchoice == 1) || (playerChoice == 3 && AIchoice == 2) {
		fmt.Println("\nTHE Player WON THIS ROUND - DAMAGE: ", damRet)
		NewAIHealth -= damRet
		damRet = 0
		return damRet, NewPHealth, NewAIHealth
	}
	fmt.Print("\n\n\nSomething bad happened\n\n\n")
	return damRet, NewPHealth, NewAIHealth
}
