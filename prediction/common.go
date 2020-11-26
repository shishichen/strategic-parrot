package prediction

import (
	"sort"

	"github.com/shishichen/strategic-parrot/base"
)

// level represents a set of holes with the same score, relative to a known board.
type level struct {
	score base.Score
	holes [][]base.Card
}

func (l *level) getScore() base.Score {
	return l.score
}

func (l *level) getHoles() [][]base.Card {
	return l.holes
}

// evaluate evaluates all possible holes chosen from the given set of cards and sorts them into levels based on their
// score when combined with the given board. The cards should not overlap with the board (but do not have to constitute
// a full deck). There must be 3 to 5 cards in the board.
func evaluate(cards, board []base.Card) []level {
	holes := base.GetCombinations(cards, 2)
	hand := make([]base.Card, len(board)+2)
	copy(hand, board)

	ranking := make(map[base.Score][][]base.Card)
	for _, hole := range holes {
		copy(hand[len(board):], hole)
		score, _ := base.GetScore(hand)
		ranking[score] = append(ranking[score], hole)
	}

	scores := []base.Score{}
	for s := range ranking {
		scores = append(scores, s)
	}
	sort.Slice(scores, func(x, y int) bool { return scores[x] > scores[y] })
	result := make([]level, len(scores))
	for i, s := range scores {
		result[i].score = s
		result[i].holes = ranking[s]
	}

	return result
}
