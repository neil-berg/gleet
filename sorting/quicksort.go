package sorting

// Partition picks a pivot, swaps items less than the pivot in the left subarray
// and then swaps the pivot into the correct position
func Partition(arr []int, lo, hi int) int {
	pivot := arr[hi] // select the rightmost element as the pivot
	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			// Pile all items less than the pivot to the left subarray
			arr[i], arr[lo] = arr[lo], arr[i]
			lo++
		}
	}

	// Finally, shift the pivot to its correct index
	arr[lo], arr[hi] = arr[hi], arr[lo]

	return lo
}

// Quicksort implements the quick sort algorithm
func Quicksort(arr []int, lo, hi int) {
	if lo > hi {
		return
	}

	newLo := Partition(arr, lo, hi)
	Quicksort(arr, lo, newLo-1)
	Quicksort(arr, newLo+1, hi)
}
