package easy

import (
	"strconv"
	"strings"
)

// PalindromeNumber checks whether a number is a palindrome.
func PalindromeNumber(n int) bool {
	if n < 0 || n%10 == 0 {
		return false
	}

	s := strconv.Itoa(n)
	chars := strings.Split(s, "")

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		if chars[i] != chars[j] {
			return false
		}
	}
	return true
}
