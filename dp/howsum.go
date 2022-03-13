package dp

import "fmt"

// HowSumCache is a cache of visited nodes and their corresponding values
type HowSumCache map[int][]int

// HowSum returns any possible way to form the sum of target given an array
func HowSum(target int, arr []int, cache HowSumCache) []int {
	cachedValue, exists := cache[target]
	if exists {
		fmt.Println("using cache")
		return cachedValue
	}

	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	for _, v := range arr {
		remainder := target - v
		result := HowSum(remainder, arr, cache)
		if result != nil {
			cache[target] = append(result, v)
			return cache[target]
		}
	}

	cache[target] = nil
	return cache[target]
}
