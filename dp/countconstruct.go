package dp

import (
	"strings"
)

type countConstructCache map[string]int

// CountConstruct returns the numbers of ways to construct the target with
// a given list of words.
func CountConstruct(target string, words []string, cache countConstructCache) int {
	cachedValue, exists := cache[target]
	if exists {
		return cachedValue
	}

	if len(target) == 0 {
		return 1
	}

	total := 0

	for _, word := range words {
		if strings.HasPrefix(target, word) {
			remainder := strings.TrimPrefix(target, word)
			numWaysForRest := CountConstruct(remainder, words, cache)
			total += numWaysForRest
		}
	}

	cache[target] = total
	return total
}
