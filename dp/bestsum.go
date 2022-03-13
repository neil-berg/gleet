package dp

// BestSumCache is a map of visited nodes and their values
type BestSumCache map[int][]int

// BestSum determines the smallest way to sum to the target.
func BestSum(target int, arr []int, cache BestSumCache) []int {
	cachedValue, exists := cache[target]
	if exists {
		return cachedValue
	}

	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	var shortest []int

	for _, v := range arr {
		remainder := target - v
		result := BestSum(remainder, arr, cache)
		if result != nil {
			result = append(result, v)
			if len(shortest) == 0 || len(result) < len(shortest) {
				shortest = result
			}
		}
	}

	cache[target] = shortest
	return shortest
}
