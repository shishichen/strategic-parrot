package prediction

import (
	"fmt"

	"github.com/shishichen/strategic-parrot/base"
)

// GetCurrentOrder returns, given a board of 3 to 5 cards, all other possible coexisting holes sorted into levels
// from best to worst, where holes in better levels will beat holes in worse levels and holes in the same level will tie,
// assuming no other cards are dealt to the board, as well as the rank of this hole relative to the ordering.
// TODO: add number of opponents
func GetCurrentOrder(hole []base.Card, board []base.Card) ([][][]base.Card, int64, int64, int64, int64, error) {
	if len(hole) != 2 {
		return nil, 0, 0, 0, 0, fmt.Errorf("current order can only be returned for holes of 2 cards")
	}
	if len(board) < 3 || len(board) > 5 {
		return nil, 0, 0, 0, 0, fmt.Errorf("current order can only be returned for boards of 3 to 5 cards")
	}

	deck := base.NewDeck()
	deck.Remove(hole)
	deck.Remove(board)
	levels := evaluate(deck.GetCards(), board)

	hand := make([]base.Card, len(hole)+len(board))
	copy(hand, hole)
	copy(hand[len(hole):], board)
	score, _ := base.GetScore(hand)

	result := make([][][]base.Card, len(levels))
	better, same, worse, total := int64(0), int64(0), int64(0), int64(0)
	for i, level := range levels {
		result[i] = level.getHoles()

		size := int64(len(level.getHoles()))
		total += size
		if level.getScore() > score {
			better += size
		} else if level.getScore() == score {
			same += size
		} else {
			worse += size
		}
	}
	return result, better, same, worse, total, nil
}
