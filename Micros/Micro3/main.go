package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		var input int
		fmt.Println("GAME START\n-----------\n")
		var ranInt int = rand.Intn(2)
		var saysInt int = rand.Intn(101)
		if ranInt == 1 {
			fmt.Print("Simon Says: ")
		}
		fmt.Print(saysInt, "\nEnter Input: ")
		fmt.Scan(&input)
		if ranInt == 1 && (input == saysInt) {
			fmt.Println("Correct, do it again!\n")
		} else if ranInt == 0 && (input != saysInt) {
			fmt.Println("Correct, do it again!\n")
		} else {
			fmt.Println("Wrong, Game over!\n")
			break
		}
	}

}
