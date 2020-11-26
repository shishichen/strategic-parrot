package base_test

import (
	"sort"
	"testing"

	"github.com/shishichen/strategic-parrot/base"
)

func TestKey(t *testing.T) {
	tests := []struct {
		name  string
		cards []base.Card
		want  base.Key
	}{
		{"empty", []base.Card{}, base.Key(0x0)},
		{"set of 2", []base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart)}, base.Key(0xD3C3)},
		{"ignores order", []base.Card{base.NewCard(base.King, base.Heart), base.NewCard(base.Ace, base.Heart)}, base.Key(0xD3C3)},
		{"set of 5", []base.Card{base.NewCard(base.Ten, base.Heart), base.NewCard(base.Jack, base.Heart),
			base.NewCard(base.Queen, base.Heart), base.NewCard(base.King, base.Heart),
			base.NewCard(base.Ace, base.Heart)}, base.Key(0xD3C3B3A393)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := base.GetKey(tt.cards)
			if err != nil || key != tt.want {
				t.Errorf("GetKey() = %x, want %x", key, tt.want)
			}
			if inverse := base.ParseKey(key); !equivalent(inverse, tt.cards) {
				t.Errorf("ParseKey() = %v, want %v", inverse, tt.cards)
			}
		})
	}
}

func equivalent(x, y []base.Card) bool {
	if len(x) != len(y) {
		return false
	}
	sort.Slice(x, func(a, b int) bool { return x[a] < x[b] })
	sort.Slice(y, func(a, b int) bool { return y[a] < y[b] })
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
