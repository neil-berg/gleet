package easy

// LongestSubstrNoRepeat finds the length of the longest substring without
// repeating characters.
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
func LongestSubstrNoRepeat(s string) int {
	seen := make(map[string]int)

	left := 0 // left position in the scan, can be updated when repeat chars are found
	max := 0  // global max length of substring

	for i, r := range s {
		seenIdx, exists := seen[string(r)]

		if !exists { // Char has not been seen yet
			window := i - left + 1
			if window > max {
				max = window
			}
		} else { // Repeat char
			if seenIdx >= left {
				// Repeat is within window. Shift left pointer to the index + 1 of
				// the seen character
				left = seenIdx + 1
			} else {
				// Repeat chat outside the window. Update max count.
				window := i - left + 1
				if window > max {
					max = window
				}
			}
		}

		// Update seen map each iteration
		seen[string(r)] = i
	}
	return max
}
