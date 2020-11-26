package base

import (
	"fmt"
	"sort"
)

// Key is a hand's key
type Key uint64

// GetKey returns an arbitrary key representing the unordered set of up to 8 cards
func GetKey(cards []Card) (Key, error) {
	if len(cards) > 8 {
		return 0, fmt.Errorf("key can only be generated for up to 8 cards")
	}

	sorted := make([]Card, len(cards))
	copy(sorted, cards)
	sort.Slice(sorted, func(x, y int) bool { return sorted[x] < sorted[y] })

	// Each 8 bits is a card, from lowest order bits to highest, from smallest card to largest
	result := uint64(0)
	for i := 0; i < len(sorted); i++ {
		result |= (uint64(sorted[i]) << (i * 8))
	}
	return Key(result), nil
}

// ParseKey parses a key into an unordered set of cards
func ParseKey(key Key) []Card {
	result := []Card{}
	for i := 0; i < 8; i++ {
		c := (key >> (i * 8)) & 0xff
		if c == 0 {
			break
		}
		result = append(result, Card(c))
	}
	return result
}
