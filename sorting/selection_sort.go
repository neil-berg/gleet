package sorting

// SelectionSort uses the selection sort algorithm to sort an array of ints
func SelectionSort(a []int) []int {
	for i := range a {
		minIdx := i

		// Find min idx in the remaining array
		for j, m := range a[i+1:] {
			if m < a[minIdx] {
				minIdx = i + j + 1
			}
		}

		// Swap min with current element
		a[i], a[minIdx] = a[minIdx], a[i]
	}
	return a
}
