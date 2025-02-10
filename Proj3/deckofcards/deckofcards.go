package deckofcards

import (
	"fmt"
	"math/rand"
)

// globals for suits and values
// go lets us declare these even if we don't use them jokerface.png
var suits []string = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
var values []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

// Card represents a single playing card
type Card struct {
	Suit  string
	Value string
}

// CardDeck represents a deck of playing cards
type CardDeck struct {
	Cards []Card
}

// NewDeck initializes a new deck of cards in standard order
func NewDeck() *CardDeck {
	deck := &CardDeck{Cards: make([]Card, 0, 52)}
	for i := 0; i < 4; i++ {
		for u := 0; u < 13; u++ {
			deck.Cards = append(deck.Cards, Card{Suit: suits[i], Value: values[u]})
		}
	}
	return deck
}

// Shuffle randomizes the order of the cards in the deck
func (d *CardDeck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) { d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i] })
}

// Contains checks if the deck contains a specific card
func (d *CardDeck) Contains(card Card) bool {
	for i := 0; i < len(d.Cards); i++ {
		if card == d.Cards[i] {
			return true
		}
	}
	return false
}

// DrawTop removes and returns the top card from the deck
func (d *CardDeck) DrawTop() Card {
	if len(d.Cards) == 0 {
		fmt.Print("Error, returning basic")
		return Card{}
	}
	var retCard Card = d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return retCard
}

// DrawBottom removes and returns the bottom card from the deck
func (d *CardDeck) DrawBottom() Card {
	if len(d.Cards) == 0 {
		fmt.Print("Error, returning basic")
		return Card{}
	}
	var retCard Card = d.Cards[0]
	d.Cards = d.Cards[1:]
	return retCard
}

// DrawRandom removes and returns a card from a random position in the deck
func (d *CardDeck) DrawRandom() Card {
	if len(d.Cards) == 0 {
		fmt.Print("Error, returning basic")
		return Card{}
	}
	var randomInt int = rand.Intn(len(d.Cards))
	var retCard Card = d.Cards[randomInt]
	d.Cards = append(d.Cards[:randomInt], d.Cards[randomInt+1:]...)
	return retCard
}

// CardToTop places a card on top of the deck
func (d *CardDeck) CardToTop(card Card) {

}

// CardToBottom places a card on the bottom of the deck
func (d *CardDeck) CardToBottom(card Card) {

}

// CardToRandom places a card at a random position in the deck
func (d *CardDeck) CardToRandom(card Card) {

}

// CardsLeft returns the number of cards left in the deck
func (d *CardDeck) CardsLeft() int {
	var length int = len(d.Cards)
	return length
}
