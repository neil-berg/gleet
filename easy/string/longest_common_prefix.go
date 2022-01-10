package easy

// LongestCommonPrefix finds the longest common prefix in a slice of strings.
func LongestCommonPrefix(arr []string) string {
	longest := ""

	var recur func(i int, arr []string)
	recur = func(i int, arr []string) {
		var char string

		for _, word := range arr {
			if i >= len(word) {
				return
			}

			if char == "" {
				char = string(word[i])
			} else if char != string(word[i]) {
				return
			}

		}
		longest += char

		recur(i+1, arr)
	}

	recur(0, arr)
	return longest
}
