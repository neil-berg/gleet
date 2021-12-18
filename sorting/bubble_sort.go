package sorting

// BubbleSort uses bubble sort to sort an array
func BubbleSort(a []int, n int) []int {
	if n == 1 {
		return a
	}

	// Each pass results in correctly setting the last element
	for i := range a[0 : len(a)-1] {
		if a[i+1] < a[i] {
			a[i], a[i+1] = a[i+1], a[i]
		}
	}

	// Recur on sub-array containing 0 through 2nd to last elements
	return BubbleSort(a, n-1)
}
