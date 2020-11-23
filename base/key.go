package base

import (
	"fmt"
	"sort"
)

// GetKey returns an arbitrary key representing the unordered set of up to 8 cards
// The input may be mutated
func GetKey(cards []Card) (uint64, error) {
	if len(cards) > 8 {
		return 0, fmt.Errorf("key can only be generated for up to 8 cards")
	}

	// Each 8 bits is a card, from lowest order bits to highest, from smallest card to largest
	sort.Slice(cards, func(x, y int) bool { return cards[x] < cards[y] })
	result := uint64(0)
	for i := 0; i < len(cards); i++ {
		result |= (uint64(cards[i]) << (i * 8))
	}
	return result, nil
}

// ParseKey parses a key into an unordered set of cards
func ParseKey(key uint64) []Card {
	result := []Card{}
	for i := 0; i < 8; i++ {
		c := (key >> (i * 8)) & 0xFF
		if c == 0 {
			break
		}
		result = append(result, Card(c))
	}
	return result
}
