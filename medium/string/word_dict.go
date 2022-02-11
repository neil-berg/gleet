package medium

import (
	"regexp"
)

// WordDict returns true if s can be segmented into a space-separated sequence
// of one or more dictionary words in wordDict,
func WordDict(s string, wordDict []string) bool {
	count := 0

	for _, word := range wordDict {
		re := regexp.MustCompile(word)
		locs := re.FindAllStringSubmatchIndex(s, -1)
		for _, loc := range locs {
			count += loc[1] - loc[0]
		}

		if count == len(s) {
			return true
		}
	}

	return count == len(s)
}
