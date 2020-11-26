package base

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck is a deck of cards
type Deck struct {
	cards []Card
}

// NewDeck returns a new unshuffled deck
func NewDeck() *Deck {
	deck := &Deck{}
	for _, rank := range []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace} {
		for _, suit := range []Suit{Club, Diamond, Heart, Spade} {
			deck.cards = append(deck.cards, NewCard(rank, suit))
		}
	}
	return deck
}

// Shuffle shuffles the deck
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(x, y int) { d.cards[x], d.cards[y] = d.cards[y], d.cards[x] })
}

// GetCards returns the remaining cards in the deck
func (d *Deck) GetCards() []Card {
	return d.cards
}

// Next deals one card from the deck
func (d *Deck) Next() (Card, error) {
	if len(d.cards) == 0 {
		return 0, fmt.Errorf("deck is empty")
	}
	c := d.cards[0]
	d.cards = d.cards[1:]
	return c, nil
}

// Remove removes a set of cards from the deck
func (d *Deck) Remove(remove []Card) {
	for _, r := range remove {
		for i, c := range d.cards {
			if c == r {
				copy(d.cards[i:], d.cards[i+1:])
				d.cards = d.cards[:len(d.cards)-1]
				break
			}
		}
	}
}
