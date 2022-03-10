package easy

// UniqueOccurrences counts the number of occurences of each integer in the array
// and returns true if there are no duplicate counts of the occurrences.
func UniqueOccurrences(arr []int) bool {
	occurences := make(map[int]int)

	for _, v := range arr {
		value, ok := occurences[v]
		if !ok {
			occurences[v] = 1
		} else {
			occurences[v] = value + 1
		}
	}

	var seen []int
	for _, v := range occurences {
		for _, v2 := range seen {
			if v == v2 {
				return false
			}
		}
		seen = append(seen, v)
	}

	return true
}
