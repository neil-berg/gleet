package easy

import "strings"

// LengthLastWord finds the length of the last word in a string s.
func LengthLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
