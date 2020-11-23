package base_test

import (
	"reflect"
	"testing"

	"github.com/shishichen/strategic-parrot/base"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		name  string
		cards []base.Card
		k     int
		want  [][]base.Card
	}{
		{"2 choose 2",
			[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart)},
			2,
			[][]base.Card{
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart)},
			},
		},
		{"3 choose 2",
			[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart),
				base.NewCard(base.Queen, base.Heart)},
			2,
			[][]base.Card{
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart)},
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.Queen, base.Heart)},
				[]base.Card{base.NewCard(base.King, base.Heart), base.NewCard(base.Queen, base.Heart)},
			},
		},
		{"5 choose 3",
			[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart),
				base.NewCard(base.Queen, base.Heart), base.NewCard(base.Jack, base.Heart),
				base.NewCard(base.Ten, base.Heart)},
			3,
			[][]base.Card{
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart),
					base.NewCard(base.Queen, base.Heart)},
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart),
					base.NewCard(base.Jack, base.Heart)},
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.King, base.Heart),
					base.NewCard(base.Ten, base.Heart)},
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.Queen, base.Heart),
					base.NewCard(base.Jack, base.Heart)},
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.Queen, base.Heart),
					base.NewCard(base.Ten, base.Heart)},
				[]base.Card{base.NewCard(base.Ace, base.Heart), base.NewCard(base.Jack, base.Heart),
					base.NewCard(base.Ten, base.Heart)},
				[]base.Card{base.NewCard(base.King, base.Heart), base.NewCard(base.Queen, base.Heart),
					base.NewCard(base.Jack, base.Heart)},
				[]base.Card{base.NewCard(base.King, base.Heart), base.NewCard(base.Queen, base.Heart),
					base.NewCard(base.Ten, base.Heart)},
				[]base.Card{base.NewCard(base.King, base.Heart), base.NewCard(base.Jack, base.Heart),
					base.NewCard(base.Ten, base.Heart)},
				[]base.Card{base.NewCard(base.Queen, base.Heart), base.NewCard(base.Jack, base.Heart),
					base.NewCard(base.Ten, base.Heart)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base.GetCombinations(tt.cards, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCombinations2(b *testing.B) {
	deck := base.NewDeck()
	for i := 0; i < b.N; i++ {
		base.GetCombinations(deck.GetCards(), 2)
	}
}

func BenchmarkCombinations5(b *testing.B) {
	deck := base.NewDeck()
	for i := 0; i < b.N; i++ {
		base.GetCombinations(deck.GetCards(), 5)
	}
}

func BenchmarkCombinations7(b *testing.B) {
	deck := base.NewDeck()
	for i := 0; i < b.N; i++ {
		base.GetCombinations(deck.GetCards(), 7)
	}
}
