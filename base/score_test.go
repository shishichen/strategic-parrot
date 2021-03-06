package base_test

import (
	"reflect"
	"testing"

	"github.com/shishichen/strategic-parrot/base"
)

func TestScore(t *testing.T) {
	tests := []struct {
		name  string
		cards []base.Card
		want  base.Score
	}{
		{"straight flush - (A,D) (K,D) (Q,D) (J,D) (T,D)",
			[]base.Card{base.NewCard(base.Ten, base.Diamond), base.NewCard(base.Ace, base.Diamond),
				base.NewCard(base.Queen, base.Diamond), base.NewCard(base.King, base.Diamond),
				base.NewCard(base.Jack, base.Diamond)},
			base.Score(0x9dcba9)},
		{"straight flush - (A,D) (Q,S) (J,D) (T,D) (9,D) (8,D) (7,D)",
			[]base.Card{base.NewCard(base.Seven, base.Diamond), base.NewCard(base.Ace, base.Diamond),
				base.NewCard(base.Eight, base.Diamond), base.NewCard(base.Ten, base.Diamond),
				base.NewCard(base.Ace, base.Spade), base.NewCard(base.Jack, base.Diamond),
				base.NewCard(base.Nine, base.Diamond)},
			base.Score(0x9a9876)},
		{"four of a kind - (A,S) (A,C) (A,H) (A,D) (K,C) (Q,H) (Q,D)",
			[]base.Card{base.NewCard(base.Ace, base.Spade), base.NewCard(base.Ace, base.Club),
				base.NewCard(base.Ace, base.Heart), base.NewCard(base.Queen, base.Heart),
				base.NewCard(base.Queen, base.Diamond), base.NewCard(base.King, base.Club),
				base.NewCard(base.Ace, base.Diamond)},
			base.Score(0x8ddddc)},
		{"four of a kind - (8,H) (8,D) (2,S) (2,C) (2,H) (2,D)",
			[]base.Card{base.NewCard(base.Two, base.Spade), base.NewCard(base.Eight, base.Heart),
				base.NewCard(base.Two, base.Club), base.NewCard(base.Two, base.Heart),
				base.NewCard(base.Two, base.Diamond), base.NewCard(base.Eight, base.Diamond)},
			base.Score(0x811117)},
		{"full house - (K,H) (K,C) (K,D) (2,S) (2,C) (2,H)",
			[]base.Card{base.NewCard(base.Two, base.Spade), base.NewCard(base.Two, base.Club),
				base.NewCard(base.King, base.Heart), base.NewCard(base.King, base.Club),
				base.NewCard(base.King, base.Diamond), base.NewCard(base.Two, base.Heart)},
			base.Score(0x7ccc11)},
		{"full house - (K,C) (Q,H) (5,H) (5,D) (3,S) (3,C) (3,H)",
			[]base.Card{base.NewCard(base.Three, base.Spade), base.NewCard(base.Three, base.Club),
				base.NewCard(base.Five, base.Heart), base.NewCard(base.King, base.Club),
				base.NewCard(base.Five, base.Diamond), base.NewCard(base.Three, base.Heart),
				base.NewCard(base.Queen, base.Heart)},
			base.Score(0x722244)},
		{"full house - (J,C) (J,H) (5,H) (5,D) (3,S) (3,C) (3,H)",
			[]base.Card{base.NewCard(base.Three, base.Spade), base.NewCard(base.Three, base.Club),
				base.NewCard(base.Five, base.Heart), base.NewCard(base.Jack, base.Club),
				base.NewCard(base.Five, base.Diamond), base.NewCard(base.Three, base.Heart),
				base.NewCard(base.Jack, base.Heart)},
			base.Score(0x7222aa)},
		{"flush - (K,S) (Q,S) (8,S) (7,S) (6,S) (5,S) (2,S)",
			[]base.Card{base.NewCard(base.Seven, base.Spade), base.NewCard(base.Eight, base.Spade),
				base.NewCard(base.Queen, base.Spade), base.NewCard(base.Two, base.Spade),
				base.NewCard(base.Five, base.Spade), base.NewCard(base.King, base.Spade),
				base.NewCard(base.Six, base.Spade)},
			base.Score(0x6cb765)},
		{"flush - (J,C) (T,D) (9,D) (8,D) (7,D) (6,S) (3,D)",
			[]base.Card{base.NewCard(base.Seven, base.Diamond), base.NewCard(base.Three, base.Diamond),
				base.NewCard(base.Eight, base.Diamond), base.NewCard(base.Ten, base.Diamond),
				base.NewCard(base.Six, base.Spade), base.NewCard(base.Jack, base.Club),
				base.NewCard(base.Nine, base.Diamond)},
			base.Score(0x698762)},
		{"straight - (T,S) (T,D) (9,H) (8,H) (7,C) (7,D) (6,C)",
			[]base.Card{base.NewCard(base.Ten, base.Spade), base.NewCard(base.Six, base.Club),
				base.NewCard(base.Seven, base.Diamond), base.NewCard(base.Nine, base.Heart),
				base.NewCard(base.Eight, base.Heart), base.NewCard(base.Seven, base.Club),
				base.NewCard(base.Ten, base.Diamond)},
			base.Score(0x598765)},
		{"straight - (A,S) (9,H) (8,H) (5,C) (4,D) (3,C) (2,D)",
			[]base.Card{base.NewCard(base.Ace, base.Spade), base.NewCard(base.Three, base.Club),
				base.NewCard(base.Two, base.Diamond), base.NewCard(base.Eight, base.Heart),
				base.NewCard(base.Nine, base.Heart), base.NewCard(base.Five, base.Club),
				base.NewCard(base.Four, base.Diamond)},
			base.Score(0x54321d)},
		{"straight - (Q,C) (T,S) (9,C) (8,D) (7,D) (6,S) (3,H)",
			[]base.Card{base.NewCard(base.Seven, base.Diamond), base.NewCard(base.Three, base.Heart),
				base.NewCard(base.Eight, base.Diamond), base.NewCard(base.Ten, base.Spade),
				base.NewCard(base.Six, base.Spade), base.NewCard(base.Queen, base.Club),
				base.NewCard(base.Nine, base.Club)},
			base.Score(0x598765)},
		{"three of a kind - (J,S) (9,D) (4,C) (4,H) (4,D)",
			[]base.Card{base.NewCard(base.Jack, base.Spade), base.NewCard(base.Four, base.Club),
				base.NewCard(base.Four, base.Heart), base.NewCard(base.Four, base.Diamond),
				base.NewCard(base.Nine, base.Diamond)},
			base.Score(0x4333a8)},
		{"three of a kind - (J,S) (9,D) (9,C) (9,H) (4,D) (2,C)",
			[]base.Card{base.NewCard(base.Two, base.Club), base.NewCard(base.Jack, base.Spade),
				base.NewCard(base.Nine, base.Club), base.NewCard(base.Nine, base.Heart),
				base.NewCard(base.Four, base.Diamond), base.NewCard(base.Nine, base.Diamond)},
			base.Score(0x4888a3)},
		{"two pair - (Q,S) (Q,H) (8,H) (8,D) (7,C) (4,C) (4,D)",
			[]base.Card{base.NewCard(base.Queen, base.Spade), base.NewCard(base.Four, base.Club),
				base.NewCard(base.Eight, base.Heart), base.NewCard(base.Seven, base.Club),
				base.NewCard(base.Four, base.Diamond), base.NewCard(base.Queen, base.Heart),
				base.NewCard(base.Eight, base.Diamond)},
			base.Score(0x3bb776)},
		{"pair - (K,C) (Q,D) (8,C) (8,H) (7,S) (6,D) (2,H)",
			[]base.Card{base.NewCard(base.Two, base.Heart), base.NewCard(base.Seven, base.Spade),
				base.NewCard(base.Eight, base.Club), base.NewCard(base.Queen, base.Diamond),
				base.NewCard(base.Eight, base.Heart), base.NewCard(base.King, base.Club),
				base.NewCard(base.Six, base.Diamond)},
			base.Score(0x277cb6)},
		{"high card - (K,C) (Q,D) (8,C) (7,S) (6,D) (5,H) (2,H)",
			[]base.Card{base.NewCard(base.Seven, base.Spade), base.NewCard(base.Eight, base.Club),
				base.NewCard(base.Queen, base.Diamond), base.NewCard(base.Two, base.Heart),
				base.NewCard(base.Five, base.Heart), base.NewCard(base.King, base.Club),
				base.NewCard(base.Six, base.Diamond)},
			base.Score(0x1cb765)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := base.GetScore(tt.cards); err != nil || got != tt.want {
				t.Errorf("GetScore() = %x, want %x", got, tt.want)
			}
		})
	}
}

func TestScoreDistribution(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		count    map[uint64]int // total number of hands by rank
		distinct map[uint64]int // number of distinct hands by rank
	}{
		{"distributions of 5 card hands", 5,
			map[uint64]int{
				1: 1302540,
				2: 1098240,
				3: 123552,
				4: 54912,
				5: 10200,
				6: 5108,
				7: 3744,
				8: 624,
				9: 40,
			},
			map[uint64]int{
				1: 1277,
				2: 2860,
				3: 858,
				4: 858,
				5: 10,
				6: 1277,
				7: 156,
				8: 156,
				9: 10,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deck := base.NewDeck()
			deck.Shuffle()
			combinations := base.GetCombinations(deck.GetCards(), tt.k)

			count := make(map[uint64]int)
			distinct := make(map[uint64]map[base.Score]bool)
			for i := uint64(1); i <= 9; i++ {
				distinct[i] = make(map[base.Score]bool)
			}
			for _, c := range combinations {
				score, _ := base.GetScore(c)
				ranking := uint64(score >> 20)
				count[ranking]++
				distinct[ranking][score] = true
			}
			if !reflect.DeepEqual(count, tt.count) {
				t.Errorf("GetScore() count = %v, want %v", count, tt.count)
			}
			for k, v := range distinct {
				if len(v) != tt.distinct[k] {
					t.Errorf("GetScore() distinct for rank %v = %v, want %v", k, len(v), tt.distinct[k])
				}
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		base.GetScore([]base.Card{base.NewCard(base.Seven, base.Diamond), base.NewCard(base.Three, base.Heart),
			base.NewCard(base.Eight, base.Diamond), base.NewCard(base.Ten, base.Spade),
			base.NewCard(base.Six, base.Spade), base.NewCard(base.Queen, base.Club),
			base.NewCard(base.Nine, base.Club)})
	}
}
