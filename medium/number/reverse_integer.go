package medium

import "math"

// ReverseInteger reverses an integer.
// 123 => 321
// -123 => -321
// 120 => 21
func ReverseInteger(n int) int {
	isNegative := n < 0

	var rev int

	for n != 0 {
		rev = rev*10 + n%10
		n = n / 10
	}

	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	if isNegative {
		rev = rev * 1
	}
	return rev
}
