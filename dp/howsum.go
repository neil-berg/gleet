package dp

// HowSum returns any possible way to form the sum of target given an array
func HowSum(target int, arr []int) []int {
	if target == 0 {
		return []int{}
	}

	if target < 0 {
		return nil
	}

	for _, v := range arr {
		remainder := target - v
		result := HowSum(remainder, arr)
		if result != nil {
			return append(result, v)
		}
	}

	return nil
}
