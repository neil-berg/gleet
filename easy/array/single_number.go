package easy

// SingleNumber finds the number that only appears once in an array of numbers
// where all others appear twice.
func SingleNumber(arr []int) int {
	seen := make(map[int]bool)

	for _, n := range arr {
		_, exists := seen[n]
		if exists {
			delete(seen, n)
		} else {
			seen[n] = true
		}
	}

	var result int
	for k := range seen {
		result = k
	}
	return result
}
