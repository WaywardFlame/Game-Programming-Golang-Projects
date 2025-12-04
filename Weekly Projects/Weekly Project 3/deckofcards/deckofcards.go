package deckofcards

import "math/rand/v2"

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
	deck := CardDeck{Cards: make([]Card, 0, 52)}
	for _, suit := range suits {
		for _, value := range values {
			deck.Cards = append(deck.Cards, Card{Suit: suit, Value: value})
		}
	}
	return &deck
}

// Shuffle randomizes the order of the cards in the deck
func (d *CardDeck) Shuffle() {
	var temp Card
	for i := 0; i < len(d.Cards); i++ {
		k := rand.IntN(52)
		temp = d.Cards[i]
		d.Cards[i] = d.Cards[k]
		d.Cards[k] = temp
	}
}

// Contains checks if the deck contains a specific card
func (d *CardDeck) Contains(card Card) bool {
	for _, elem := range d.Cards {
		if elem == card {
			return true
		}
	}
	return false
}

// DrawTop removes and returns the top card from the deck
func (d *CardDeck) DrawTop() Card {
	var topCard Card = d.Cards[0]
	d.Cards = d.Cards[1:]
	return topCard
}

// DrawBottom removes and returns the bottom card from the deck
func (d *CardDeck) DrawBottom() Card {
	var bottomCard Card = d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return bottomCard
}

// DrawRandom removes and returns a card from a random position in the deck
func (d *CardDeck) DrawRandom() Card {
	num := rand.IntN(len(d.Cards))
	var randomCard Card = d.Cards[num]
	d.Cards = append(d.Cards[:num], d.Cards[num+1:]...)
	return randomCard
}

// CardToTop places a card on top of the deck
func (d *CardDeck) CardToTop(card Card) {
	d.Cards = append([]Card{card}, d.Cards...)
}

// CardToBottom places a card on the bottom of the deck
func (d *CardDeck) CardToBottom(card Card) {
	d.Cards = append(d.Cards, card)
}

// CardToRandom places a card at a random position in the deck
func (d *CardDeck) CardToRandom(card Card) {
	num := rand.IntN(len(d.Cards))
	tempSlice := d.Cards[num:]
	d.Cards = append(d.Cards[:num], []Card{card}...)
	d.Cards = append(d.Cards, tempSlice...)
}

// CardsLeft returns the number of cards left in the deck
func (d *CardDeck) CardsLeft() int {
	return len(d.Cards)
}
