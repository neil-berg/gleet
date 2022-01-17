package easy

// SearchInsertPosition return the index if the target is found. If not, return
// the index where it would be if it were inserted in order.
func SearchInsertPosition(arr []int, target int) int {
	// Implementing binary search, time complexity of O(log n)
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if target > arr[mid] {
			start = mid + 1
		} else if target < arr[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}

	return start
}
