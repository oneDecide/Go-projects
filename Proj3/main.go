package main

import (
	"Project3/deckofcards"
	"fmt"
)

func main() {

	MyTests() //run your custom test code

	points := 0

	//Create a new deck of cards
	deck := deckofcards.NewDeck()
	expectedDeckSize := 52
	if deck.CardsLeft() == expectedDeckSize {
		points += 5
		fmt.Println("Deck initialized correctly. Points awarded :) ", points)
	} else {
		fmt.Println("Deck initialization failed. ")
	}

	// Shuffle the deck
	deck.Shuffle()
	if isSufficientlyShuffled(deck, 10) {
		fmt.Println("Deck shuffled correctly. Points awarded :) ", points)
		points += 5
	} else {
		fmt.Println("Deck not shuffled correctly.")
	}

	// Draw the top card
	topCard := deck.DrawTop()
	if topCard != (deckofcards.Card{}) && deck.CardsLeft() == expectedDeckSize-1 {
		points += 5
		fmt.Println("Top card drawn correctly. Points awarded :) ", points)
	} else {
		fmt.Println("Drawing top card failed.")
	}

	// Draw the bottom card
	bottomCard := deck.DrawBottom()
	if bottomCard != (deckofcards.Card{}) && deck.CardsLeft() == expectedDeckSize-2 {
		points += 5
		fmt.Println("Bottom card drawn correctly. Points awarded :) ", points)
	} else {
		fmt.Println("Drawing bottom card failed.")
	}

	// Draw a random card
	randomCard := deck.DrawRandom()
	if randomCard != (deckofcards.Card{}) && deck.CardsLeft() == expectedDeckSize-3 {
		points += 5
		fmt.Println("Random card drawn correctly. Points awarded :) ", points)
	} else {
		fmt.Println("Drawing random card failed.")
	}

	// Add the top card back to the top of the deck
	deck.CardToTop(topCard)
	if deck.Contains(topCard) && deck.CardsLeft() == expectedDeckSize-2 {
		points += 5
		fmt.Println("Top card added back to top correctly. Points awarded :) ", points)
	} else {
		fmt.Println("Adding top card back to top failed.")
	}

	// Add the bottom card back to the bottom of the deck
	deck.CardToBottom(bottomCard)
	if deck.Contains(bottomCard) && deck.CardsLeft() == expectedDeckSize-1 {
		points += 5
		fmt.Println("Bottom card added back to bottom correctly. Points awarded :) ", points)
	} else {
		fmt.Println("Adding bottom card back to bottom failed.")
	}

	// Add a random card to a random position in the deck
	deck.CardToRandom(randomCard)
	if deck.Contains(randomCard) && deck.CardsLeft() == expectedDeckSize {
		points += 5
		fmt.Println("Random card added back to a random position correctly. Points awarded :) ", points)
	} else {
		fmt.Println("Adding random card back to random position failed.")
	}

	// Check if a specific card is in the deck
	cardToCheck := deckofcards.Card{Suit: "Hearts", Value: "A"}
	contains := deck.Contains(cardToCheck)
	if contains {
		points += 5
		fmt.Println("Deck contains the specified card. Points awarded :) ", points)
	} else {
		fmt.Println("Deck does not contain the specified card.")
	}

	// Print the number of cards left in the deck
	if deck.CardsLeft() == expectedDeckSize {
		points += 5
		fmt.Println("Number of cards left in the deck is correct. Points awarded :) ", points)
	} else {
		fmt.Println("Number of cards left in the deck is incorrect.")
	}
	fmt.Println("\nTotal Points awarded: ", points)
}

// extremely naive test because I'm lazy
func isSufficientlyShuffled(deck *deckofcards.CardDeck, minOutOfOrder int) bool {
	originalDeck := deckofcards.NewDeck()
	outOfOrderCount := 0

	for i, card := range deck.Cards {
		if card != originalDeck.Cards[i] {
			outOfOrderCount++
		}

		if outOfOrderCount >= minOutOfOrder {
			return true
		}
	}
	return false
}
