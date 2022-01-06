package easy

// TwoSum returns indices of the two numbers such that they add up to target
// given an array of integers nums and an integer target
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].
func TwoSum(nums []int, target int) []int {
	// Track value:index as we traverse in one loop
	numsToIdx := make(map[int]int)

	for i, n := range nums {
		// what remainder do we need to meet the target?
		remainder := target - n

		if remainderIdx, ok := numsToIdx[remainder]; ok {
			return []int{remainderIdx, i}
		}

		// No pair found, store this num and move on to the next value
		numsToIdx[n] = i
	}

	return []int{0, 0}
}
