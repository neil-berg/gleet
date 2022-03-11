package dp

import (
	"strconv"
)

// GridTravellerCache is the cache of visited nodes and corresponding values
type GridTravellerCache map[string]int

// GridTraveller determines the number of paths to travel a grid of m rows
// by n columns if you only travel down or right.
func GridTraveller(m, n int, cache GridTravellerCache) int {
	// Construct a cache key of "m,n"
	key := strconv.Itoa(m) + "," + strconv.Itoa(n)
	// Note that GridTraveller(m,n) = GridTraveller(n,m), so we can check for
	// reverse cache keys as well.
	revKey := strconv.Itoa(n) + "," + strconv.Itoa(m)

	value, exists := cache[key]
	if exists {
		return value
	}

	value, exists = cache[revKey]
	if exists {
		return value
	}

	if m == 1 && n == 1 {
		// At the end box
		return 1
	}

	if m == 0 || n == 0 {
		// Out of bounds
		return 0
	}

	// Add to the cache
	cache[key] = GridTraveller(m-1, n, cache) + GridTraveller(m, n-1, cache)

	return cache[key]
}
