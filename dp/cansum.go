package dp

// CanSumCache is a cache of visited nodes and corresponding values
type CanSumCache map[int]bool

// CanSum determines if an array of numbers can be summed to the target.
func CanSum(arr []int, target int, cache CanSumCache) bool {
	cachedValue, exists := cache[target]
	if exists {
		return cachedValue
	}

	if target == 0 {
		return true
	}

	if target < 0 {
		return false
	}

	for _, v := range arr {
		remainder := target - v
		if CanSum(arr, remainder, cache) {
			cache[remainder] = true
			return true
		}
	}

	cache[target] = false
	return false
}
