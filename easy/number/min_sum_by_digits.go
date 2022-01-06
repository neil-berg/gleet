package easy

import (
	"sort"
)

// Given an array of digits (values are from 0 to 9), find the minimum possible
// sum of two numbers formed from digits of the array. All digits of given array
// must be used to form the two numbers.
// Example:
// Input: [6, 8, 4, 5, 2, 3]
// Output: 604
// The minimum sum is formed by numbers
// 358 and 246

// MinSumByDigits finds the minimum sum of two numbers formed from digits of the
// input array.
func MinSumByDigits(arr []int) int {
	sort.Ints(arr)

	var n1 int
	var n2 int

	for i, v := range arr {
		// Alternate digit
		if i%2 == 0 {
			// Even indices to n1, odds to n2
			n1 = n1*10 + v
		} else {
			n2 = n2*10 + v
		}
	}

	return n1 + n2
}
