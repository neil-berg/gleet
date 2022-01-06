package easy

import (
	"strconv"
	"strings"
)

// absInt returns the absolute value of an int
func absInt(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// ReverseInt reverses the digits of an interger.
func ReverseInt(n int) int {
	isNegative := n < 0
	s := strconv.Itoa(absInt(n))
	chars := strings.Split(s, "")
	// Reverse the characters in-place
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	i, _ := strconv.Atoi(strings.Join(chars, ""))
	if isNegative {
		i = i * -1
	}
	return i
}
