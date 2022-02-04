package easy

// ValidParens checks if a string has valid parentheses.
func ValidParens(s string) bool {
	var stack []string

	if len(s)%2 != 0 {
		return false
	}

	for _, r := range s {
		// The concept is to append the right-pair whenever a left paren is found.
		// Then when a right paren is found, check that it matches top element.
		if string(r) == "(" {
			stack = append(stack, ")")
		} else if string(r) == "{" {
			stack = append(stack, "}")
		} else if string(r) == "[" {
			stack = append(stack, "]")
		} else {
			// Right-paren
			if len(stack) == 0 || stack[len(stack)-1] != string(r) {
				return false
			}
			// Match, pop last element
			stack = stack[:len(stack)-1]
		}
	}
	return true
}
