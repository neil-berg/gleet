package easy

// DuplicatesOnce returns any duplicate items in the array once in the output array.
// Ex: [1,1,2,2,2] => [1,2]
func DuplicatesOnce(arr []int) []int {
	seen := make(map[int]bool)

	var dupes []int

	for _, n := range arr {
		inDupes, exists := seen[n]
		if !exists {
			// First occurrence, mark as seen and not in dupes
			seen[n] = false
		} else if !inDupes {
			dupes = append(dupes, n)
			seen[n] = true

		}
	}

	return dupes
}
