package medium

// newSet is an internal helper to convert an array into a set
func newSet(a []int) map[int]bool {
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}
	return m
}

// LongestConsecutiveSequence takes an unsorted array nums and returns the length
// of the longest consecutive elements sequence.
// Ex: [100, 4, 200, 7, 5, 6] => 4
func LongestConsecutiveSequence(nums []int) int {
	set := newSet(nums)

	longestSeq := 1
	for _, v := range nums {
		currSeq := 1
		currNum := v
		for set[currNum+1] == true {
			currSeq++
			currNum++
		}
		if currSeq > longestSeq {
			longestSeq = currSeq
		}
	}
	return longestSeq
}
