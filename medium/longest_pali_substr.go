package medium

// isPalindrome checks whether the string is a palindrome
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// LongestPaliSubstr returns the longest palindromic substring in s.
func LongestPaliSubstr(s string) string {
	// Record past char:index pairs
	seen := make(map[string][]int)

	longest := ""

	for i, r := range s {
		indices, exists := seen[string(r)]
		if exists {
			// Repeat char, first see if i - seen_index > longest and if so, check if
			// it's a palindrome
			for _, index := range indices {
				if i-index > len(longest) && isPalindrome(s[index:i+1]) {
					longest = s[index : i+1]
				}
			}
		}

		// Add to seen map
		seen[string(r)] = append(seen[string(r)], i)
	}

	return longest
}
