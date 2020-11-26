package base_test

import (
	"testing"

	"github.com/shishichen/strategic-parrot/base"
)

func TestNext(t *testing.T) {
	tests := []struct {
		name    string
		dealt   int
		success bool
	}{
		{"full", 0, true},
		{"middle", 16, true},
		{"empty", 52, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := base.NewDeck()
			deck.Shuffle()
			seen := make(map[base.Card]bool)
			for i := 0; i < tt.dealt; i++ {
				if c, err := deck.Next(); err != nil || seen[c] {
					t.Errorf("Next() returned duplicate card %v", c)
				} else {
					seen[c] = true
				}
			}
			if remaining := len(deck.GetCards()); remaining != (52 - tt.dealt) {
				t.Errorf("GetCards() length = %v, want length %v", remaining, 52-tt.dealt)
			}
			if _, err := deck.Next(); (err == nil) != tt.success {
				t.Errorf("NewCard() error = %v, want success %v", err, tt.success)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name      string
		remove    []base.Card
		remaining int
	}{
		{"nothing", []base.Card{}, 52},
		{"something", []base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart),
			base.NewCard(base.Queen, base.Heart), base.NewCard(base.Jack, base.Heart),
			base.NewCard(base.Ten, base.Heart)}, 47},
		{"everything", base.NewDeck().GetCards(), 0},
		{"duplicates", []base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.Ace, base.Heart)}, 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := base.NewDeck()
			deck.Shuffle()
			deck.Remove(tt.remove)
			if len(deck.GetCards()) != tt.remaining {
				t.Errorf("remaining deck length = %v, want length %v", len(deck.GetCards()), tt.remaining)
			}
			remaining := make(map[base.Card]bool)
			for _, c := range deck.GetCards() {
				remaining[c] = true
			}
			for _, c := range tt.remove {
				if remaining[c] {
					t.Errorf("card %v was removed but still appears in deck", c)
				}
			}
		})
	}
}
