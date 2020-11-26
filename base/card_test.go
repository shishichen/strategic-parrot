package base_test

import (
	"testing"

	"github.com/shishichen/strategic-parrot/base"
)

func TestNewCard(t *testing.T) {
	tests := []struct {
		name string
		rank base.Rank
		suit base.Suit
		want base.Card
	}{
		{"(4,C)", base.Four, base.Club, base.Card(0x31)},
		{"(7,D)", base.Seven, base.Diamond, base.Card(0x62)},
		{"(2,H)", base.Two, base.Heart, base.Card(0x13)},
		{"(Q,S)", base.Queen, base.Spade, base.Card(0xb4)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base.NewCard(tt.rank, tt.suit); got != tt.want {
				t.Errorf("NewCard() = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		card base.Card
		want string
	}{
		{"(4,C)", base.NewCard(base.Four, base.Club), "(4,C)"},
		{"(7,D)", base.NewCard(base.Seven, base.Diamond), "(7,D)"},
		{"(2,H)", base.NewCard(base.Two, base.Heart), "(2,H)"},
		{"(Q,S)", base.NewCard(base.Queen, base.Spade), "(Q,S)"},
		{"invalid", base.Card(0xff), "(invalid)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
