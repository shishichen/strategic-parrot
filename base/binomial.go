package base

// GetCombinations returns all unordered combinations of size k from the given cards
// Optimization note: returning a flat slice makes the (52 choose 7) benchmark considerably faster, but not
// sure we care right now
func GetCombinations(cards []Card, k int) [][]Card {
	n := len(cards)
	if k < 0 || k > n {
		return [][]Card{}
	}
	result := make([][]Card, numCombinations(n, k))
	generateCombinations([]Card{}, cards, k, result, 0)
	return result
}

// Requires: 0 <= k <= n
func numCombinations(n, k int) int {
	if n-k < k {
		k = n - k
	}
	numerator := 1
	denominator := 1
	for i := 1; i <= k; i++ {
		numerator *= n + 1 - i
		denominator *= i
	}
	return numerator / denominator
}

// generateCombinations generates combinations with the fixed prefix, choosing k from the remaining
// Requires: 0 <= k <= n
func generateCombinations(prefix, remaining []Card, k int, result [][]Card, index int) int {
	if k == 0 {
		result[index] = make([]Card, len(prefix))
		copy(result[index], prefix)
		return index + 1
	}
	for i := 0; i < len(remaining)-(k-1); i++ {
		prefix = append(prefix, remaining[i])
		index = generateCombinations(prefix, remaining[i+1:], k-1, result, index)
		prefix = prefix[:len(prefix)-1]
	}
	return index
}
