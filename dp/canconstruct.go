package dp

import (
	"strings"
)

// CanConstructCache is a map of visited nodes and their values
type CanConstructCache map[string]bool

// CanConstruct determines if target can be constructued with words.
func CanConstruct(target string, words []string, cache CanConstructCache) bool {
	cachedValue, exists := cache[target]
	if exists {
		return cachedValue
	}

	if len(target) == 0 {
		return true
	}

	for _, word := range words {
		// Only branch with prefixes
		if strings.HasPrefix(target, word) {
			remainder := strings.TrimPrefix(target, word)
			if CanConstruct(remainder, words, cache) {
				cache[target] = true
				return true
			}
		}

	}

	cache[target] = false
	return false
}
