package sorting

// Merge combines left and right arrays, which are sorted
func Merge(left, right []int) []int {
	size := len(left) + len(right)

	// New array of merged items from left and right
	merged := make([]int, size)

	// i and j are positions in the left and right arrays
	i, j := 0, 0

	// Iterate until all items in left and right are merged
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			// Exhausted left only, select next item in right
			merged[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			// Exhausted right, select next item in left
			merged[k] = left[i]
			i++
		} else if left[i] < right[j] {
			// Left and right are not exhausted, next left item is smaller than next right
			merged[k] = left[i]
			i++
		} else {
			// Left and right are not exhausted, next right item is smaller than next left
			merged[k] = right[j]
			j++
		}
	}
	return merged
}

// MergeSort recursively divides an array by half until len 1 and then calls
// Merge to sort and combine them in a sorted array
func MergeSort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	mid := len(a) / 2
	left, right := a[:mid], a[mid:]
	return Merge(MergeSort(left), MergeSort(right))
}
