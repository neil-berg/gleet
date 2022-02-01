package medium

import "strings"

// ReverseWords returns the words reversed in a string, accounting for whitespace.
// Input: s = "  a good   example"
// Output: "example good a"
func ReverseWords(s string) string {
	var out string

	words := strings.Fields(s)
	for i := len(words) - 1; i >= 0; i-- {
		if i == 0 {
			out += words[i]
		} else {
			out += words[i] + " "
		}
	}

	return out
}
