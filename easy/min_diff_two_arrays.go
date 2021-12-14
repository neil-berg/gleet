package easy

import (
	"math"
	"sort"
)

// MinDiffTwoArrays computes the min difference between to unsorted arrays.
// A: [5,1,9,12]
// B: [30,8,3,25]
// Output: 1 (9,8)
func MinDiffTwoArrays(a, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	min := math.MaxInt8

	ai := 0
	bi := 0

	for ai < len(a) && bi < len(b) {
		if int(math.Abs(float64(a[ai]-b[bi]))) < min {
			min = int(math.Abs(float64(a[ai] - b[bi])))
		}

		// Update index pointers - min value advances
		if a[ai] < b[bi] {
			ai++
		} else {
			bi++
		}
	}
	return min
}
