package heap

// Heap is the shape of a heap
type Heap struct {
	arr []int
}

// MaxHeapify converts a 0-based array into a max-heap
func (h *Heap) MaxHeapify(i int) {
	// For node at index i:
	leftChildIdx := 2*i + 1
	rightChildIdx := 2*i + 2

	// Track the index of the largest value between current parent node (i-th
	// index) and its children
	largestIdx := i

	if leftChildIdx < len(h.arr) && h.arr[leftChildIdx] > h.arr[i] {
		largestIdx = leftChildIdx
	}
	if rightChildIdx < len(h.arr) && h.arr[rightChildIdx] > h.arr[i] {
		largestIdx = rightChildIdx
	}

	if largestIdx != i {
		// Swap the parent with its larger child node
		h.arr[i], h.arr[largestIdx] = h.arr[largestIdx], h.arr[i]
		// Recursively heapify (downward) possibly affected subtrees after the swap
		h.MaxHeapify(largestIdx)
	}
}

// CreateMaxHeap reverse walks an array, starting at the last non-leaf node index
// (i.e. the first parent node index) and turns it into a max-heap.
func (h *Heap) CreateMaxHeap() {
	// Last non-leaf node index in a heap is N/2-1
	for i := len(h.arr)/2 - 1; i >= 0; i-- {
		h.MaxHeapify(i)
	}
}
