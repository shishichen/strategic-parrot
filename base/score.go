package base

import (
	"fmt"
	"sort"
)

// GetScore returns an arbitrary but ordered number representing the value of the best poker hand,
// out of 5, 6, or 7 cards
// The input may be mutated
func GetScore(cards []Card) (uint64, error) {
	if len(cards) < 5 || len(cards) > 7 {
		return 0, fmt.Errorf("only a set of 5, 6, or 7 cards can be scored")
	}

	ranking, significant := getRanking(cards)
	// First 40 bits are unusued. Next 4 bits contain the hand ranking. Next 5 sets of 4 bits contain
	// the ranks of the 5 significant cards, in order from most to least significant.
	score := uint64(0)
	score |= (uint64(ranking) << 20)
	for i := 0; i < 5; i++ {
		score |= (uint64(significant[i]) << (16 - i*4))
	}
	return score, nil
}

type ranking int

const (
	highCard      ranking = iota + 1 // significant ranks: cards from highest to lowest
	pair                             // significant ranks: pair, remaining cards from highest to lowest
	twoPair                          // significant ranks: highest pair, lowest pair, remaining highest card
	threeOfAKind                     // significant ranks: three of a kind, remaining cards from highest to lowest
	straight                         // significant ranks: straight from highest to lowest
	flush                            // significant ranks: flush cards from highest to lowest
	fullHouse                        // significant ranks: three of a kind, pair
	fourOfAKind                      // significant ranks: four of a kind, highest remaining card
	straightFlush                    // significant ranks: straight flush cards from highest to lowest
)

// Returns the best ranking and the ranks of the 5 significant cards, in order from most to least
func getRanking(cards []Card) (ranking, []Rank) {
	significant := make([]Rank, 5)
	// Sort from largest to smallest
	sort.Slice(cards, func(x, y int) bool { return cards[x] > cards[y] })

	// First, check for a straight flush
	haveFlush, flushCards := hasFlush(cards)
	if haveFlush {
		haveStraight, straightFlushCards := hasStraight(flushCards)
		if haveStraight {
			for i := 0; i < 5; i++ {
				significant[i] = straightFlushCards[i].GetRank()
			}
			return straightFlush, significant
		}
	}

	// Next, four of a kind
	m, n, mRank, nRank := hasSomeOfAKind(cards)
	if m == 4 {
		for i := 0; i < 4; i++ {
			significant[i] = mRank
		}
		significant[4] = nRank
		return fourOfAKind, significant
	}

	// Next, full house
	if m == 3 && n == 2 {
		for i := 0; i < 3; i++ {
			significant[i] = mRank
		}
		for i := 3; i < 5; i++ {
			significant[i] = nRank
		}
		return fullHouse, significant
	}

	// Next, flush
	if haveFlush {
		for i := 0; i < 5; i++ {
			significant[i] = flushCards[i].GetRank()
		}
		return flush, significant
	}

	// Next, straight
	haveStraight, straightCards := hasStraight(cards)
	if haveStraight {
		for i := 0; i < 5; i++ {
			significant[i] = straightCards[i].GetRank()
		}
		return straight, significant
	}

	// Next, three of a kind
	if m == 3 {
		for i := 0; i < 3; i++ {
			significant[i] = mRank
		}
		significant[3] = nRank
		for _, c := range cards {
			r := c.GetRank()
			if r != mRank && r != nRank {
				significant[4] = r
				break
			}
		}
		return threeOfAKind, significant
	}

	// Next, two pair
	if m == 2 && n == 2 {
		for i := 0; i < 2; i++ {
			significant[i] = mRank
		}
		for i := 2; i < 4; i++ {
			significant[i] = nRank
		}
		for _, c := range cards {
			r := c.GetRank()
			if r != mRank && r != nRank {
				significant[4] = r
				break
			}
		}
		return twoPair, significant
	}

	// Next, pair
	if m == 2 {
		for i := 0; i < 2; i++ {
			significant[i] = mRank
		}
		significant[2] = nRank
		i := 3
		for _, c := range cards {
			r := c.GetRank()
			if r != mRank && r != nRank {
				significant[i] = r
				i++
			}
			if i > 4 {
				break
			}
		}
		return pair, significant
	}

	// Finally, if we have nothing else, we just have a high card
	for i := 0; i < 5; i++ {
		significant[i] = cards[i].GetRank()
	}
	return highCard, significant
}

func hasSomeOfAKind(cards []Card) (m int, n int, mRank Rank, nRank Rank) {
	counts := make(map[Rank]int)
	for _, c := range cards {
		counts[c.GetRank()]++
	}
	for r, count := range counts {
		if count > m || (count == m && r > mRank) {
			m, mRank = count, r
		}
	}
	for r, count := range counts {
		if r == mRank {
			continue
		}
		if count > 5-m {
			count = 5 - m
		}
		if count > n || (count == n && r > nRank) {
			n, nRank = count, r
		}
	}
	return m, n, mRank, nRank
}

func hasFlush(cards []Card) (bool, []Card) {
	counts := make(map[Suit][]Card)
	for _, c := range cards {
		counts[c.GetSuit()] = append(counts[c.GetSuit()], c)
	}
	for _, ranks := range counts {
		if len(ranks) >= 5 {
			return true, ranks // should already be in order
		}
	}
	return false, []Card{}
}

func isConsecutive(x Card, y Card) bool {
	r := x.GetRank()
	if r == Two {
		return y.GetRank() == Ace
	}
	return y.GetRank() == (r - 1)
}

func hasStraight(cards []Card) (bool, []Card) {
	unique := []Card{cards[0]}
	for i := 1; i < len(cards); i++ {
		if cards[i].GetRank() != cards[i-1].GetRank() {
			unique = append(unique, cards[i])
		}
	}
	if cards[0].GetRank() == Ace {
		unique = append(unique, cards[0])
	}
	if len(unique) < 5 {
		return false, []Card{}
	}

	start := 0
	for i := 1; i < len(unique); i++ {
		if !isConsecutive(unique[i-1], unique[i]) {
			// If we already have a straight without the next card, return now
			if i-start >= 5 {
				return true, unique[start:i]
			}
			// Try again, starting from the next card, unless we don't have enough cards left
			if len(unique)-i < 5 {
				return false, []Card{}
			}
			start = i
		}
	}
	// If we made it outside the loop, it means the last few cards were consecutive
	return true, unique[start:]
}
