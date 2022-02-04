package medium

func reverseString(s string) string {
	var res string
	for _, r := range s {
		res = string(r) + res
	}
	return res
}

// ReverseBetweenParens Reverse the strings in each pair of matching parentheses,
// starting from the innermost one.
// Input: s = "(u(love)i)"
// Output: "iloveu"
// Input: s = "(ed(et(oc))el)"
// Output: "leetcode"
func ReverseBetweenParens(s string) string {
	// i and j are pointers starting in the middle of the string, with i advancing
	// leftward and j advancing rightward between left and right parens, respectively.
	i, j := len(s)/2, len(s)/2

	// Keep track of i, j whenever they hit a paren, i => ( and j => )
	iPrev, jPrev := len(s)/2, len(s)/2

	res := ""

	for i >= 0 && j <= len(s)-1 {
		if string(s[i]) != "(" {
			i--
		}

		if string(s[j]) != ")" {
			j++
		}

		if string(s[i]) == "(" && string(s[j]) == ")" {
			leftPart := s[i+1 : iPrev]
			// For the initial reversal only, we use jPrev as the initial index.
			// All others need to use jPrev + 1 since jPrev is on the paren now.
			var rightPart string
			if res == "" {
				rightPart = s[jPrev:j]
			} else {
				rightPart = s[jPrev+1 : j]
			}
			res = reverseString(leftPart + res + rightPart)
			// Record where i and j are before moving to the next chunk
			iPrev = i
			jPrev = j
			// Advance i and j away from their current parentheses
			i--
			j++
		}
	}

	return res
}
