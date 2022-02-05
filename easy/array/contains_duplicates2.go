package easy

// absDiff is a helper to return the absolute difference between a and b.
func absDiff(a, b int) int {
	diff := a - b
	if diff < 0 {
		diff = diff * -1
	}
	return diff
}

// ContainsDuplicates2 returns true if there are two distinct indices i and j in
// the array such that nums[i] == nums[j] and abs(i - j) <= k.
// Input: nums = [1,2,3,1], k = 3
// Output: true
func ContainsDuplicates2(arr []int, k int) bool {
	seen := make(map[int]int)

	for i, n := range arr {
		j, exists := seen[n]
		if exists && absDiff(i, j) <= k {
			return true
		}
		// Always update seen[n] with the most recent index of this value's occurence
		seen[n] = i
	}

	return false
}
