package prediction

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/shishichen/strategic-parrot/base"
)

// GetFutureOutcomes returns, given a hole and board of 3 to 5 cards, the probability that the hole will win, tie,
// and lose at the end of the game, assumming random subsequent cards.
// TODO: add number of opponents
func GetFutureOutcomes(hole []base.Card, board []base.Card) (float64, float64, float64, error) {
	if len(hole) != 2 {
		return 0, 0, 0, fmt.Errorf("future outcomes can only be predicted for holes with 2 cards")
	}
	if len(board) == 0 {
		return 0, 0, 0, fmt.Errorf("call GetInitialOutcomes to get outcomes for an empty board ")
	}
	if len(board) < 3 || len(board) > 5 {
		return 0, 0, 0, fmt.Errorf("future outcomes can only be predicted for boards with 3 to 5 cards")
	}

	deck := base.NewDeck()
	deck.Remove(hole)
	deck.Remove(board)
	subsequent := base.GetCombinations(deck.GetCards(), 5-len(board))

	n := runtime.NumCPU()
	var wg sync.WaitGroup
	better, same, worse, total := make([]int64, n), make([]int64, n), make([]int64, n), make([]int64, n)
	for id := 0; id < n; id++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			futureBoard := make([]base.Card, 5)
			copy(futureBoard, board)

			futureHand := make([]base.Card, len(hole)+5)
			copy(futureHand, hole)
			copy(futureHand[len(hole):], board)

			lower := id * len(subsequent) / n
			upper := (id + 1) * len(subsequent) / n
			for i := lower; i < upper; i++ {
				deck := base.NewDeck()
				deck.Remove(hole)
				deck.Remove(board)
				deck.Remove(subsequent[i])
				copy(futureBoard[len(board):], subsequent[i])
				levels := evaluate(deck.GetCards(), futureBoard)

				copy(futureHand[len(hole)+len(board):], subsequent[i])
				score, _ := base.GetScore(futureHand)

				for _, level := range levels {
					size := int64(len(level.getHoles()))
					total[id] += size
					if level.getScore() > score {
						better[id] += size
					} else if level.getScore() == score {
						same[id] += size
					} else {
						worse[id] += size
					}
				}
			}
		}(id)
	}
	wg.Wait()

	betterTotal, sameTotal, worseTotal, totalTotal := int64(0), int64(0), int64(0), int64(0)
	for id := 0; id < n; id++ {
		betterTotal += better[id]
		sameTotal += same[id]
		worseTotal += worse[id]
		totalTotal += total[id]
	}
	return float64(worseTotal) / float64(totalTotal), float64(sameTotal) / float64(totalTotal), float64(betterTotal) / float64(totalTotal), nil
}
