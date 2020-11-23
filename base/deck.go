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

// NewDeck returns a new shuffled deck
func NewDeck() *Deck {
	deck := &Deck{}
	for _, rank := range []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace} {
		for _, suit := range []Suit{Club, Diamond, Heart, Spade} {
			deck.cards = append(deck.cards, NewCard(rank, suit))
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(x, y int) { deck.cards[x], deck.cards[y] = deck.cards[y], deck.cards[x] })
	return deck
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
