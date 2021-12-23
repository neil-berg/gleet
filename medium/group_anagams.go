package medium

// GroupAnagrams groups anagrams in a given an array of strings strs.
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
func GroupAnagrams(strs []string) [][]string {
	// Map of ASCII-sum (key) to array of words (value)
	m := make(map[int][]string)

	for _, str := range strs {
		sum := 0
		for _, n := range str {
			sum += int(n)
		}
		m[sum] = append(m[sum], str)
	}

	output := make([][]string, len(m))
	i := 0
	for _, v := range m {
		output[i] = append(output[i], v...)
		i++
	}

	return output
}
