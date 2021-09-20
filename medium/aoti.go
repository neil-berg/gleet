package medium

import (
	"regexp"
	"strconv"
	"strings"
)

const maxInt = 2 ^ 31

// Atoi converts a string to an integer
func Atoi(s string) int32 {
	// Not valid if does not start with digit or -+digit
	goodPrefix := regexp.MustCompile(`^(\s+)?(\-|\+)?\d`).MatchString(s)
	if !goodPrefix {
		return 0
	}

	match := regexp.MustCompile(`((\-|\+)?[0-9]+)`).FindString(s)
	i, err := strconv.ParseInt(match, 10, 32)
	if err != nil {
		leadingChar := strings.Split(match, "")[0]
		if leadingChar == "-" {
			return maxInt * -1
		}
		return maxInt
	}
	return int32(i)
}
